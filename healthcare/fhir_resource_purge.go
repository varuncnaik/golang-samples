// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package snippets

// [START healthcare_delete_resource_purge]
import (
	"context"
	"fmt"
	"io"

	healthcare "google.golang.org/api/healthcare/v1"
)

// purgeFHIRResource purges an FHIR resources.
func purgeFHIRResource(w io.Writer, projectID, location, datasetID, fhirStoreID, resourceType, fhirResourceID string) error {
	ctx := context.Background()

	healthcareService, err := healthcare.NewService(ctx)
	if err != nil {
		return fmt.Errorf("healthcare.NewService: %w", err)
	}

	fhirService := healthcareService.Projects.Locations.Datasets.FhirStores.Fhir

	name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/fhirStores/%s/fhir/%s/%s", projectID, location, datasetID, fhirStoreID, resourceType, fhirResourceID)

	if _, err := fhirService.ResourcePurge(name).Do(); err != nil {
		return fmt.Errorf("ResourcePurge: %w", err)
	}

	fmt.Fprintf(w, "Resource Purged: %q", name)

	return nil
}

// [END healthcare_delete_resource_purge]
