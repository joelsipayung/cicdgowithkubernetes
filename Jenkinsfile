pipeline {

  environment {
    registry = "joelsipayung/cicdjenkinskubego"
    dockerImage = ""
  }

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/joelsipayung/cicdgowothkubernetes.git'
      }
    }

    stage('Build image') {
      steps{
        script {
          dockerImage = docker.build (registry)
        }
      }
    }

    stage('Push Image') {
      steps{
        script {
          docker.withRegistry( 'https://registry.hub.docker.com', 'joelsipayung' ) {
            dockerImage.push("${env.BUILD_NUMBER}")
            dockerImages.push("latest")
          }
        }
      }
    }


   stage('Deploy App') {
      steps {
        script {
          kubeconfig(credentialsId: 'mykubeconfig', serverUrl: 'https://usw1.kubesail.com') {
            sh 'kubectl apply -f mytask.yaml'
          }
        }
      }
    }


  }

}
