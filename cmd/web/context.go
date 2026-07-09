package main

// We need our type for context key to avoid name collisions with
// 3rd party packages
type contextKey string

const isAuthenticatedContextKey = contextKey("isAuthenticated")
