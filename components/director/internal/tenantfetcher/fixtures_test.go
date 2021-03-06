package tenantfetcher_test

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/kyma-incubator/compass/components/director/internal/tenantfetcher"
)

func fixEvent(id, name string, fieldMapping tenantfetcher.TenantFieldMapping) []byte {
	eventData := fmt.Sprintf(`{"%s":"%s","%s":"%s"}`, fieldMapping.IDField, id, fieldMapping.NameField, name)

	return []byte(fmt.Sprintf(`{
		"id":        %s,
		"eventData": %s,
	}`, fixID(), eventData))
}

func fixEventWithDiscriminator(id, name, discriminator string, fieldMapping tenantfetcher.TenantFieldMapping) []byte {
	discriminatorData := ""
	if fieldMapping.DiscriminatorField != "" {
		discriminatorData = fmt.Sprintf(`"%s": "%s",`, fieldMapping.DiscriminatorField, discriminator)
	}

	eventData := fmt.Sprintf(`{"%s":"%s",%s"%s":"%s"}`, fieldMapping.IDField, id, discriminatorData, fieldMapping.NameField, name)

	return []byte(fmt.Sprintf(`{
		"id":        %s,
		"eventData": %s,
	}`, fixID(), eventData))
}

func fixBusinessTenantMappingInput(name, externalTenant, provider string) model.BusinessTenantMappingInput {
	return model.BusinessTenantMappingInput{
		Name:           name,
		ExternalTenant: externalTenant,
		Provider:       provider,
	}
}

func fixTenantEventsResponse(events []byte, total, pages int) tenantfetcher.TenantEventsResponse {
	return tenantfetcher.TenantEventsResponse(fmt.Sprintf(`{
		"events":       %s,
		"total": %d,
		"pages":   %d,
	}`, string(events), total, pages))
}

func fixID() string {
	return uuid.New().String()
}
