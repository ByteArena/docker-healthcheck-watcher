#!/bin/sh
#https://docs.microsoft.com/en-us/azure/devops/pipelines/process/variables?view=azure-devops&tabs=yaml%2Cbatch#set-in-script

# Global, should not change
TIME_STAMP=$"`date "+%Y%m%d-%H%M"`"
echo '##vso[task.setvariable variable=TIME_STAMP]'${TIME_STAMP}
if [[ ${COMPLETE_RELEASE_BRANCH} =~ refs/tags/(.+)$ ]]; then
    echo '##vso[task.setvariable variable=DOCKER_TAG]'${BASH_REMATCH[1]}
    echo '##vso[task.setvariable variable=GIT_TAG]'${BASH_REMATCH[1]}
else
    echo '##vso[task.setvariable variable=DOCKER_TAG]'preview-${RELEASE_BRANCH}-${TIME_STAMP}
fi

# Might change
echo '##vso[task.setvariable variable=REGISTRY]stratsys.azurecr.io'

# Per project, should change every time
echo '##vso[task.setvariable variable=REPOSITORY]docker-healthcheck-watcher'
