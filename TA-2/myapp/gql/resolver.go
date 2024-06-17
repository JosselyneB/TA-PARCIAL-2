package gql

type Resolver struct{}

func (r *Resolver) Hello() string {
    return "Hello, world!"
}
