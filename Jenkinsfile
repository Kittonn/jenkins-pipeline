pipeline {
  agent any

  environment {
    IMAGE_NAME = "echo-api"
    GITHUB_CREDENTIALS = credentials('kitton-github')
  }  

  stages {
    stage("Build") {
      steps {
        sh "docker build --tag ${IMAGE_NAME} ."
        sh "docker image ls"
      }
    }
  }
}
