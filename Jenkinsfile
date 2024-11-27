pipeline {
  agent any

  environment {
    IMAGE_NAME = "echo-api"
    GITHUB_CREDENTIALS = credentials('kitton-github')
  }  

  stages {
    stage("Build") {
      steps {
        sh "docker buildx build --push --tag ${IMAGE_NAME} --platform linux/amd64,linux/arm64 ."
      }
    }
  }
}
