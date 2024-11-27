pipeline {
  agent any

  environment {
    IMAGE_NAME = "echo-api"
    GITHUB_CREDENTIALS = credentials('kitton-github')
  }  

  stages {
    stage("Login to GitHub Container Registry") {
      steps { 
        sh "echo ${GITHUB_CREDENTIALS_PSW} | docker login ghcr.io -u ${GITHUB_CREDENTIALS_USR} --password-stdin"
      }
    }

    

    stage("Build Docker Image") {     
      steps {
        sh "docker buildx build --tag ghcr.io/${GITHUB_CREDENTIALS_USR}/${IMAGE_NAME}:${BUILD_NUMBER} --push ."
      }
    }
  }
}
