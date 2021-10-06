def SERVICE_GROUP = "scarif"
def SERVICE_NAME = "golang-hands-on"
def IMAGE_NAME = "${SERVICE_GROUP}-${SERVICE_NAME}"
def REPOSITORY_URL = "ssh://git@code.bespinglobal.com/scal/golang-hands-on.git"
def REPOSITORY_SECRET = "bespin-poc-ssh"
def SLACK_TOKEN_DEV = ""
def SLACK_TOKEN_DQA = ""

@Library("scarif-pipeline-library")
def butler = new com.bespin.scarif.JenkinsPipeline()
def label = "worker-${UUID.randomUUID().toString()}"

properties([
  buildDiscarder(logRotator(daysToKeepStr: "60", numToKeepStr: "30"))
])
podTemplate(label: label, containers: [
  containerTemplate(name: "builder", image: "opsnowtools/valve-builder:v0.2.60", command: "cat", ttyEnabled: true, alwaysPullImage: true),
  containerTemplate(name: "gobuilder", image: "opsnowtools/valve-gobuilder:v0.2.0", command: "cat", ttyEnabled: true, alwaysPullImage: true)
], volumes: [
  hostPathVolume(mountPath: "/var/run/docker.sock", hostPath: "/var/run/docker.sock"),
  hostPathVolume(mountPath: "/home/jenkins/.draft", hostPath: "/home/jenkins/.draft"),
  emptyDirVolume(mountPath: '/home/jenkins', memory: false)
], envVars: [
  envVar(key: 'HOME', value: '/home/jenkins')
]) {
   node(label) {
    stage("Prepare") {
      container("builder") {
        butler.prepare(IMAGE_NAME)
        sh "chmod 666 /var/run/docker.sock"
      }
    }

    stage("Checkout") {
      container("builder") {
        try {
          if (REPOSITORY_SECRET) {
            git(url: REPOSITORY_URL, branch: BRANCH_NAME, credentialsId: REPOSITORY_SECRET)
          } else {
            git(url: REPOSITORY_URL, branch: BRANCH_NAME)
          }
        } catch (e) {
          butler.failure(SLACK_TOKEN_DEV, "Checkout")
          throw e
        }

        butler.scan("golang")
      }
    }
    stage("Build") {
          container("gobuilder") {
            try {
              butler.golang_swagger()
              butler.golang_build()
              butler.success(SLACK_TOKEN_DEV, "Build")
            } catch (e) {
              butler.failure(SLACK_TOKEN_DEV, "Build")
              throw e
            }
          }
        }
    stage("Tests") {
        container("gobuilder") {
          try {
            butler.golang_unit_test()
          } catch (e) {
            butler.failure(SLACK_TOKEN_DEV, "Tests")
            throw e
        }
      }
    }

    stage("Code Analysis") {
        container("gobuilder") {
          try {
            butler.golang_sonar()
            butler.success(SLACK_TOKEN_DEV, "Code Analysis")
          } catch (e) {
            butler.failure(SLACK_TOKEN_DEV, "Code Analysis")
            throw e
          }
        } // container 끝나는 부분
    }
    if (BRANCH_NAME == "master") {
      stage("Build Image") {
        parallel(
          "Build Docker": {
            container("builder") {
              try {
                butler.build_image()
              } catch (e) {
                butler.failure(SLACK_TOKEN_DEV, "Build Docker")
                throw e
              }
            }
          },
          "Build Charts": {
            container("builder") {
              try {
                butler.build_chart()
              } catch (e) {
                butler.failure(SLACK_TOKEN_DEV, "Build Charts")
                throw e
              }
            }
          }
        )
      }
      stage('Prisma Cloud Scan') {
        // Scan the image
        this.VERSION = butler.get_version()
        echo "# scan version: ${VERSION}"
        prismaCloudScanImage ca: '',
          cert: '',
          dockerAddress: 'unix:///var/run/docker.sock',
          image: "harbor-devops.coruscant.opsnow.com/opsnow/${IMAGE_NAME}:${VERSION}",
          key: '',
          logLevel: 'info',
          podmanPath: '',
          project: '',
          resultsFile: 'prisma-cloud-scan-results.json',
          ignoreImageBuildTime:true
      }
      stage('Prisma Cloud Publish') {
        prismaCloudPublish resultsFilePattern: 'prisma-cloud-scan-results.json'
      }
      stage("Deploy DEV") {
        container("builder") {
          try {
            // deploy(cluster, namespace, sub_domain, profile)
            butler.deploy("dev", "${SERVICE_GROUP}-dev", "golang-hands-on", "dev", "site/okc1/values.dev.yaml")
            butler.success(SLACK_TOKEN_DEV, "Deploy DEV")
          } catch (e) {
            butler.failure(SLACK_TOKEN_DEV, "Deploy DEV")
            throw e

          }
        }
      }

      // FIXME Skip QA Process
      stage("Proceed PROD") {
        container("builder") {
          butler.proceed(SLACK_TOKEN_DQA, "Deploy PROD", "prod")
          timeout(time: 60, unit: "MINUTES") {
            input(message: "${butler.name} ${butler.version} to prod")
          }
        }
      }
      stage("Deploy PROD") {
        container("builder") {
          try {
            // deploy(cluster, namespace, sub_domain, profile)
            butler.deploy("okc1", "${SERVICE_GROUP}-prod", "golang-hands-on", "prod", "site/okc1/values.prod.yaml")
            butler.success([SLACK_TOKEN_DEV,SLACK_TOKEN_DQA], "Deploy PROD")
          } catch (e) {
            butler.failure([SLACK_TOKEN_DEV,SLACK_TOKEN_DQA], "Deploy PROD")
            throw e
          }
        }
      }
    }
  }
}


