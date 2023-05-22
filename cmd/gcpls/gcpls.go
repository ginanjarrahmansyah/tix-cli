package gcpls

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/compute/metadata"
	"github.com/spf13/cobra"
	"google.golang.org/api/compute/v1"

	"github.com/olekukonko/tablewriter"
)

type Instance struct {
	Zone      string
	Instance  string
	PrivateIP string
	PublicIP  string
}

var (
	projectID string
)

func NewCmdGCPLS() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gcpls",
		Short: "Google Cloud Platform related commands",
		Run:   runGCPLS,
	}

	cmd.Flags().StringVarP(&projectID, "project", "p", "", "Google Cloud project ID")
	cmd.MarkFlagRequired("project")

	return cmd
}

func runGCPLS(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	// if Project ID is not provied, ty to fetch it from the metadata server
	if projectID == "" {
		var err error
		projectID, err = metadata.ProjectID()
		if err != nil {
			log.Fatal("Failed to get project ID from metadata server:", err)
		}
	}

	// Create a compute service client
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal("Failed to create compute service client:", err)
	}

	// List instances in al zone of the project
	zoneList, err := computeService.Zones.List(projectID).Do()
	if err != nil {
		log.Fatal("Failed to list zones:", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Zone", "Instance", "Private IP", "Public IP"})

	instanceChan := make(chan *Instance)
	doneChan := make(chan bool)

	// launch Goroutines to fetch instances from each zone

	for _, zone := range zoneList.Items {
		go func(zone string) {
			instances, err := computeService.Instances.List(projectID, zone).Do()
			if err != nil {
				log.Fatal("Failed to retrieve instances:", err)
			}

			for _, instance := range instances.Items {
				var privateIP, publicIP string

				for _, iface := range instance.NetworkInterfaces {
					if iface.NetworkIP != "" {
						privateIP = iface.NetworkIP
					}

					if iface.AccessConfigs != nil && len(iface.AccessConfigs) > 0 {
						publicIP = iface.AccessConfigs[0].NatIP
					}
				}

				instanceInfo := &Instance{
					Zone:      zone,
					Instance:  instance.Name,
					PrivateIP: privateIP,
					PublicIP:  publicIP,
				}
				instanceChan <- instanceInfo
			}

			doneChan <- true
		}(zone.Name)
	}

	// Start a Goroutine to wait for all instances to be processed
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(len(zoneList.Items))

		go func() {
			wg.Wait()
			close(instanceChan)
		}()

		for range zoneList.Items {
			<-doneChan
			wg.Done()
		}
	}()

	// Process the received instance information concurrently
	for instance := range instanceChan {
		table.Append([]string{instance.Zone, instance.Instance, instance.PrivateIP, instance.PublicIP})
	}

	table.Render()
}
