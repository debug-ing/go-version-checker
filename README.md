# Go version checker github action

This project simple github action project for get version golang project.
My first attempts at writing a GitHub action


## User
Add this job to .yml file in workflows 

```
jobs:
  build:
    #add this  
    - name: Run Go Version Checker
      uses: debug-ing/go-version-checker@v1.0.0
```