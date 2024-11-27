pipeline {
  agent any

  environment {
    IMAGE_NAME = "echo-api"
    GITHUB_CREDENTIALS = credentials('kitton-github')
  }

  stages {
    stage("Setup Buildx") {
      steps {
        // Ensure that buildx is installed
        sh 'docker buildx create --use'
      }
    }
    
    stage("Build") {
      steps {
        sh "docker buildx build --tag ${IMAGE_NAME} ."
        sh "docker image ls"
      }
    }
  }
}
