pipeline{
  
  agent any
  
  environment {
    dockerhub=credentials('dockerhub')
    branch=env.GIT_BRANCH
  }

  stages {
    stage('docker image build') {
      steps {
        sh 'docker build -t safderun/ibu-yemek-api:build .'
      }
    }

    stage('docker image master tag'){
      when {
        environment.branch == 'master'
      }
      steps {
        sh 'docker tag safderun/ibu-yemek-api:build safderun/ibu-yemek-api:latest'
      }
    }

    stage('docker image dev tag'){
      when {
        environment.branch == 'dev'
      }
      steps {
        sh 'docker tag safderun/ibu-yemek-api:build safderun/ibu-yemek-api:dev'
      }
    }
  
    stage('docker image push') {
      steps{
        sh 'echo $dockerhub_PSW | docker login -u $dockerhub_USR --password-stdin'
        sh 'docker push safderun/ibu-yemek-api:dev'
      }
    }

    stage('deploy') {
      agen {label 'ec2'}
      steps{
          sh '/ibuYemekBotu/updateContainer.sh'
      }
    }
  }
}