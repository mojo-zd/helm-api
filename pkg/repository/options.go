package repository

type repoOption func(options *repoOptions)

// WithName ...
func WithName(name string) repoOption {
	return func(options *repoOptions) {
		options.name = name
	}
}

// WithURL repo url
func WithURL(url string) repoOption {
	return func(options *repoOptions) {
		options.url = url
	}
}

// WithUsername set username if repo need certification
func WithUsername(username string) repoOption {
	return func(options *repoOptions) {
		options.username = username
	}
}

// WithPassword set password if repo need certification
func WithPassword(password string) repoOption {
	return func(options *repoOptions) {
		options.password = password
	}
}

// WithRepoFile set repo file location
func WithRepoFile(repoFile string) repoOption {
	return func(options *repoOptions) {
		options.repoFile = repoFile
	}
}

// WithRepoCache set repo cache file location
func WithRepoCache(repoCache string) repoOption {
	return func(options *repoOptions) {
		options.repoCache = repoCache
	}
}

// WithInsecureSkipTLSverify set insecure skip tls verify
func WithInsecureSkipTLSverify(skip bool) repoOption {
	return func(options *repoOptions) {
		options.insecureSkipTLSverify = skip
	}
}
