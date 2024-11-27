pipeline {
  agent any

  environment {
    IMAGE_NAME = "echo-api"
    GITHUB_CREDENTIALS = credentials('kitton-github')
  }  

  stages {
    stage("Build Docker Image"){
      steps {
        sh "docker buildx build -t ghcr.io/${GITHUB_CREDENTIALS_USR}/${IMAGE_NAME}:${BUILD_NUMBER} ."
      }
    }


    stage("Lojgin to GitHub Container Registry") {
      steps { 
        sh "docker login ghcr.io -u ${GITHUB_CREDENTIALS_USR} -p ${GITHUB_CREDENTIALS_PSW}"
      }
    }

    stage("Push Docker Image to ghcr.io") {
      steps {
        sh "docker push ghcr.io/${GITHUB_CREDENTIALS_USR}/${IMAGE_NAME}:${BUILD_NUMBER}"
      } 
    }
  }
}
