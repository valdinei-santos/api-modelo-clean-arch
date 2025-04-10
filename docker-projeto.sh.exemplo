#!/bin/bash
# Script feita para gerar imagem docker e enviar para o servidor de HOMOLOGAÇÃO e PRODUÇÃO
PROJETO=modelos
SIGLA_PROJETO=te
APP=api-modelo-clean-arch
LANGUAGE=go
 
IMAGE_NAME=${APP}
DIR_SCRIPTS=/home/$USER/docker/${PROJETO}
DIR_BASE=/home/$USER/code/${PROJETO}
DIR_CODE=${DIR_BASE}/${LANGUAGE}
DIR_APP=${DIR_SCRIPTS}/${LANGUAGE}/${IMAGE_NAME}
DIR_DOCKERFILE=${DIR_CODE}/${IMAGE_NAME}
DIR_IMAGE_TAR=${DIR_SCRIPTS}/${LANGUAGE}/${IMAGE_NAME}/imagens
 
QTD_PORTA=1
QTD_VOLUME=1
CONTAINER_PORT=8888
CONTAINER_VOLUME=/app/logs
HOST_PORT=8801
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
 
#. ${DIR_SCRIPTS}/script-cmd-docker.sh $1

CONTAINER_NAME=${SIGLA_PROJETO}-${IMAGE_NAME}
NETWORK_NAME=bridge
#ENV_FILE=.env

IMAGE_VERSION=$1
FILENAME_IMAGE=imagem-${IMAGE_NAME}-v${IMAGE_VERSION}
FILENAME_IMAGE_TAR=${FILENAME_IMAGE}.tar
DOCKER_SERVER_PROD=192.168.1.1
DOCKER_SERVER_HOM=192.168.1.2
DOCKER_SERVER_DIR_IMAGE=/home/$USER/docker/${PROJETO}/${LANGUAGE}/${IMAGE_NAME}/imagens

opcao1() {
  IMAGE_LAST=$(docker image ls| grep ${IMAGE_NAME} | head -n 1 | awk '{print $2}' )
  echo ${IMAGE_LAST}
}

opcao2(){
  CMD="docker build -t ${IMAGE_NAME}:v${IMAGE_VERSION} ${DIR_DOCKERFILE}"
  echo "${CMD}"
  ${CMD}
}

opcao_build(){
  if [ $1 == "dev" ]; then
    DOCKERFILE="Dockerfile"
  else
    DOCKERFILE="Dockerfile-"$1
  fi
  #CMD="docker build -t ${IMAGE_NAME}:v${IMAGE_VERSION}-$1 -f ${FILENAME_DOCKERFILE_DEV} ${DIR_DOCKERFILE}"
  CMD="docker build -t ${IMAGE_NAME}:v${IMAGE_VERSION}-$1 -f ${DOCKERFILE} ${DIR_DOCKERFILE}"
  echo "${CMD}"
  ${CMD}
}

opcao3(){
  if [ "${CONTAINER_PORT}" == "" ] && [ "${CONTAINER_VOLUME}" == "" ]; then
    CMD="docker run --rm --name ${CONTAINER_NAME} --network ${NETWORK_NAME} --env-file ${ENV_FILE} -d ${IMAGE_NAME}:v${IMAGE_VERSION}"
  else 
    CMD="docker run --rm --name ${CONTAINER_NAME} -p ${HOST_PORT}:${CONTAINER_PORT} -v ${HOST_VOLUME}:${CONTAINER_VOLUME} --network ${NETWORK_NAME} --env-file ${ENV_FILE} -d ${IMAGE_NAME}:v${IMAGE_VERSION}"
  fi
  echo "${CMD}"
  ${CMD}  
}

opcao4(){
  CMD="docker stop ${CONTAINER_NAME}"
  CMD2="docker rm ${CONTAINER_NAME}"
  echo "${CMD}"
  echo "${CMD2}"
  ${CMD}
  ${CMD2}
}

opcaoR(){
  opcao4
  opcao2
  opcao3
}

opcaoX(){
  docker ps|grep ${CONTAINER_NAME}
}

opcaoL(){
  docker logs ${CONTAINER_NAME}
}

opcaoE(){
  echo -n "Digite o tipo do shell (bash/ash): "
  read tipo
  docker exec -it ${CONTAINER_NAME} $tipo
}

opcao5(){
  IMAGE_TAR_LAST=$(ls -ltr ${DIR_IMAGE_TAR}| grep ${IMAGE_NAME} | tail -n 1 | awk '{print $9}' )
  echo ${IMAGE_TAR_LAST}
}

opcao6(){
  CMD="docker save -o ${DIR_IMAGE_TAR}/${FILENAME_IMAGE_TAR} ${IMAGE_NAME}:v${IMAGE_VERSION}"
  echo "${CMD}"
  ${CMD}
  chmod g+rw ${DIR_IMAGE_TAR}/${FILENAME_IMAGE_TAR}
}

opcao_gera_tar(){
  FILENAME_IMAGE_TAR_FRONT=imagem-${IMAGE_NAME}-v${IMAGE_VERSION}-$1.tar
  CMD="docker save -o ${DIR_IMAGE_TAR}/${FILENAME_IMAGE_TAR_FRONT} ${IMAGE_NAME}:v${IMAGE_VERSION}-$1"
  echo "${CMD}"
  ${CMD}
  chmod g+rw ${DIR_IMAGE_TAR}/${FILENAME_IMAGE_TAR_FRONT}
}

opcao7(){
  scp -p ${DIR_IMAGE_TAR}/${FILENAME_IMAGE_TAR} $USER@${DOCKER_SERVER_HOM}:${DOCKER_SERVER_DIR_IMAGE}
  ssh $USER@${DOCKER_SERVER_HOM} "docker load -i ${DOCKER_SERVER_DIR_IMAGE}/${FILENAME_IMAGE_TAR}"
}

opcao8(){
  scp -p ${DIR_IMAGE_TAR}/${FILENAME_IMAGE_TAR} $USER@${DOCKER_SERVER_PROD}:${DOCKER_SERVER_DIR_IMAGE}
  ssh $USER@${DOCKER_SERVER_PROD} "docker load -i ${DOCKER_SERVER_DIR_IMAGE}/${FILENAME_IMAGE_TAR}"
}

msgFim(){
  echo ""
  echo "<Enter> Volta ao menu"
  echo "<S>     Sair"
  read -n1 x
  if [ "$x" == "S" ] || [ "$x" == "s" ]; then
    exit 1
  fi
}

while true 
do
  echo ""
  echo "Escolha a opcao:           Versao: v${IMAGE_VERSION}" 
  echo "    1) Ultima versao da imagem ${IMAGE_NAME} disponível no Docker"
  echo "    2) BUILD imagem ${IMAGE_NAME}:v${IMAGE_VERSION} "
  echo "    3) START container ${CONTAINER_NAME} com a imagem ${IMAGE_NAME}:v${IMAGE_VERSION}"
  echo "    4) STOP container ${CONTAINER_NAME}"
  echo "    R) REBUILD/RESTART container ${CONTAINER_NAME} com a imagem ${IMAGE_NAME}:v${IMAGE_VERSION}"
  echo "    L) SHOW logs container ${CONTAINER_NAME} rodando"
  echo "    X) SHOW container ${CONTAINER_NAME} rodando"
  echo "    E) EXEC -it no container ${CONTAINER_NAME}"
  echo "       5) Ultimo tar da imagem ${IMAGE_NAME} disponível"
  echo "       6) Gerar tar da imagem ${IMAGE_NAME}:v${IMAGE_VERSION} "
  echo "       7) Enviar tar da imagem ${IMAGE_NAME}:v${IMAGE_VERSION} para SERVIDOR HOMOLOGACAO (faz LOAD no docker)"
  echo "       8) Enviar tar da imagem ${IMAGE_NAME}:v${IMAGE_VERSION} para SERVIDOR PRODUCAO (faz LOAD no docker)"
  echo "    s) Sair..."
  echo -ne "--> "
  read -n1 OPCAO
  echo ""

  case $OPCAO in
    1) opcao1 ;;
    2) opcao2 ;;
    3) opcao3 ;;
    4) opcao4 ;;
    r|R) opcaoR ;;
    l|L) opcaoL ;;
    x|X) opcaoX ;;
    e|E) opcaoE ;;
    5) opcao5 ;;
    6) opcao6 ;;
    7) opcao7 ;;
    8) opcao8 ;;
    s|S)
      echo "Saindo... "
      exit 1 ;;
    *) echo "Opcao invalida!!! <ENTER> para sair"
        msgFim ;;
  esac
done
