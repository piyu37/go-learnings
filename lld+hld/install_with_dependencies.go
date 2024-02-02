package main

// You are working on a team to develop a new package manager.
// Your team has already created the following functions for you to use:

// - installPackage(packageName: string) -> (success: bool)
//   - given a package name, will attempt to download and install it
//   - will FAIL if package dependencies are not already installed on the system
//   - returns true/false based on success/fail

// - getDependencies(packageName: string) -> (packageNames: string[])
//   - given a package name, will return the names of the dependent packages

// - isInstalled(packageName: string) -> (installed: bool)
//   - given a package name, will return whether package is installed on the system

// However, many packages have dependencies which must be installed first.
// Your task is to create a new function:
// installWithDependencies(packageName: string) -> (success: bool)
//  - given a package name, will install all dependencies before installing the given package
//  - must be able to handle arbitrarily deep dependency chains
//  - must call the existing functions provided by your team (don't worry about implementing them)

// For example, given the following dependency relationship:
// A -> B,C
// B -> D,E
// In order to install package A, we would need to first install packages B,C
// But in order to install package B, we would need to first install D, E

func getDependencies(packageName string) []string {
	return []string{}
}

func installPackage(packageName string) {

}

func installWithDependencies(packageName string) bool {
	deps := getDependencies(packageName)
	for i := range deps {
		installWithDependencies(deps[i])
	}

	installPackage(packageName)

	return true
}
