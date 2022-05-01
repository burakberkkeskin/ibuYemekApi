pipeline{
  
  agent any
  
  environment {
    dockerhub=credentials('dockerhub')
  }

  stages {
    stage('docker image build') {
      steps {
        sh 'docker build -t safderun/ibu-yemek-api:build .'
        sh "docker tag safderun/ibu-yemek-api:build safderun/ibu-yemek-api:$BUILD_NUMBER"
      }
    }
  
    stage('docker image push') {
      steps{
        sh 'echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin'
        sh 'docker push safderun/ibu-yemek-api:dev'
      }
    }

    stage('deploy') {
      agent {label 'ec2'}
      steps{
          sh '/ibuYemekBotu/updateContainer.sh'
      }
    }
  }
}