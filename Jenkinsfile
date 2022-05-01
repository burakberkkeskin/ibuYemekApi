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
        sh "docker tag $IMAGE_NAME:build $IMAGE_NAME:$BUILD_NUMBER"
      }
    }
  
    stage('docker image push') {
      steps{
        sh "echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin"
        sh "docker push $IMAGE_NAME:$BUILD_NUMBER"
        sh "docker rmi $IMAGE_NAME:$BUILD_NUMBER"
      }
    }

    //Deployment
    stage('deploy') {
      agent {label 'ec2'}
      steps{
        sh "/ibuYemekBotu/updateApi.sh $BUILD_NUMBER"
      }
    }

    
  }

  post {
        success {
          mail (bcc: '', body: "Latest deploy for IBU Yemek Api was successfull!. \n Build Number: $BUILD_NUMBER", cc: '', from: 'Jenkins', replyTo: '', subject: 'IBU Yemek Api Deploy Succesfull!', to: 'safderun@proton.me')
        }
        failure {
          mail bcc: '', body: '''Latest deploy for IBU Yemek Api was failed!. 
          Build Number: $BUILD_NUMBER''', cc: '', from: 'Jenkins', replyTo: '', subject: '!!!IBU Yemek Api Deploy Failed!!!', to: 'safderun@proton.me'
        }
    }
}