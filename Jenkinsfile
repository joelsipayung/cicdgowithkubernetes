pipeline {

  agent any

  stages {

    stage('Checkout Source') {
      steps {
        git 'https://github.com/joelsipayung/cicdgowithkubernetes.git'
      }
    }

    stage('Build image') {
      steps{
        script {
          dockerImage = docker.build("joelsipayung/cicdkubewithgo:${env.BUILD_ID}")
        }
      }
    }

    stage('Push Image') {
      steps{
        script {
          docker.withRegistry( 'https://registry.hub.docker.com', 'joelsipayung' ) {
                dockerImage.push("latest")
                dockerImage.push("${env.BUILD_ID}")
          }
        }
      }
    }


   stage('Deploy App') {
      steps {
        script {
          kubeconfig(credentialsId: 'mykubeconfig', serverUrl: 'https://usw1.kubesail.com') {
            sh 'kubectl apply -f mytask.yaml'
            sh 'kubectl set image deployment/mytaskgo mytaskgo=joelsipayung/cicdkubewithgo:${env.BUILD_ID}'
          }
        }
      }
    }


  }

}
