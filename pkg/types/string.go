package types

import "github.com/DuskEagle/landscape/pkg/resource"

type StringOutput *stringOutput

type stringOutput struct {
	s string
	// Can this work??
	resource resource.Resource
	await    func(resource.Resource) string
}

func String(s string) StringOutput {
	return &stringOutput{
		s: s,
	}
}

// TODO(joel): If this works, it needs to be hidden.
func StringInternal(resource resource.Resource, await func(resource resource.Resource) string) StringOutput {
	return &stringOutput{
		resource: resource,
		await:    await,
	}
}

func (s *stringOutput) String() string {
	return s.await(s.resource)
}
