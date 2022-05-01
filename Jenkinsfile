pipeline{
  
  agent any
  
  environment {
    dockerhub=credentials('dockerhub')
    IMAGE_NAME="safderun/ibu-yemek-api"
  }

  stages {
    stage('docker image build') {
      steps {
        sh "docker build -t $IMAGE_NAME:build ."
        sh "docker tag IMAGE_NAME:build $IMAGE_NAME:$BUILD_NUMBER"
      }
    }
  
    stage('docker image push') {
      steps{
        sh "echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin"
        sh "docker push $IMAGE_NAME:$BUILD_NUMBER"
        sh "docker rmi $IMAGE_NAME:$BUILD_NUMBER"
      }
    }

    stage('deploy') {
      agent {label 'ec2'}
      steps{
          sh '/ibuYemekBotu/updateContainer.sh'
      }
    }

    post {
        success {
            echo 'This will run only if successful'
        }
        failure {
            echo 'This will run only if failed'
        }
    }
  }
}