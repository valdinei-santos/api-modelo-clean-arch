#!/bin/bash
# Script feita para gerar imagem docker e enviar para o servidor previg-api-hom e previg-api
PROJETO=autoatendimento
APP=api-modelo-clean-arch
LANGUAGE=go
 
IMAGE_NAME=${APP}
DIR_SCRIPTS=/home/previg/docker/${PROJETO}
DIR_BASE=/home/$USER/code/${PROJETO}
DIR_CODE=${DIR_BASE}/${LANGUAGE}
DIR_APP=${DIR_SCRIPTS}/${LANGUAGE}/${IMAGE_NAME}
DIR_DOCKERFILE=${DIR_CODE}/${IMAGE_NAME}/
DIR_IMAGE_TAR=${DIR_SCRIPTS}/${LANGUAGE}/${IMAGE_NAME}/imagens/
 
QTD_PORTA=1
QTD_VOLUME=1
CONTAINER_PORT=8800
CONTAINER_VOLUME=/app/logs
HOST_PORT=8812
HOST_VOLUME=${DIR_APP}/logs
ENV_FILE=${DIR_APP}/.env
 
if [ $# -eq 0 ]; then
  #IMAGE_LAST=$(docker image ls| grep ${IMAGE_NAME} | head -n 1 | awk '{print $2}' )
  echo ""
  echo "   Parametro com a versao da imagem eh obrigatorio."
  echo "   A versao deve seguir o padrao X.Y.Z"
  echo "      Exemplo: $0 0.1.0"
  echo ""
  #echo "   Ultima versao disponivel no Docker é: ${IMAGE_LAST}"
  echo "Versoes disponiveis no Docker: "
  docker image ls| grep ${IMAGE_NAME}
  echo ""
  exit 1;
fi
 
. ${DIR_SCRIPTS}/script-cmd-docker.sh $1
