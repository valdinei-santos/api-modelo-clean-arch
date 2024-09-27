#!/bin/bash
# Script feita para gerar imagem docker e colocar para rodar
APP=api-modelo-clean-arch
IMAGE_NAME=${APP}
DIR_APP=/home/$USER/code/modelos/go/${APP}
DIR_DOCKERFILE=${DIR_APP}
 
QTD_PORTA=1
QTD_VOLUME=1
CONTAINER_PORT=8888
CONTAINER_VOLUME=/app/logs
HOST_PORT=8888
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
 
CONTAINER_NAME=cont-${IMAGE_NAME}
NETWORK_NAME=bridge
IMAGE_VERSION=$1
FILENAME_IMAGE=imagem-${IMAGE_NAME}-v${IMAGE_VERSION}
FILENAME_IMAGE_TAR=${FILENAME_IMAGE}.tar

opcao1() {
  IMAGE_LAST=$(docker image ls| grep ${IMAGE_NAME} | head -n 1 | awk '{print $2}' )
  echo ${IMAGE_LAST}
}

opcao2(){
  CMD="docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} ${DIR_DOCKERFILE}"
  echo "${CMD}"
  ${CMD}
}

opcao_build(){
  DOCKERFILE="Dockerfile"
  CMD="docker build -t ${IMAGE_NAME}:${IMAGE_VERSION}-$1 -f ${DOCKERFILE} ${DIR_DOCKERFILE}"
  echo "${CMD}"
  ${CMD}
}

opcao3(){
  if [ "${CONTAINER_PORT}" == "" ] && [ "${CONTAINER_VOLUME}" == "" ]; then
    CMD="docker run --rm --name ${CONTAINER_NAME} --network ${NETWORK_NAME} --env-file ${ENV_FILE} -d ${IMAGE_NAME}:${IMAGE_VERSION}"
  else 
    CMD="docker run --rm --name ${CONTAINER_NAME} -p ${HOST_PORT}:${CONTAINER_PORT} -v ${HOST_VOLUME}:${CONTAINER_VOLUME} --network ${NETWORK_NAME} --env-file ${ENV_FILE} -d ${IMAGE_NAME}:${IMAGE_VERSION}"
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
  echo "Escolha a opcao:           Versao: ${IMAGE_VERSION}" 
  echo "    1) Ultima versao da imagem ${IMAGE_NAME} disponível no Docker"
  echo "    2) BUILD imagem ${IMAGE_NAME}:${IMAGE_VERSION} "
  echo "    3) START container ${CONTAINER_NAME} com a imagem ${IMAGE_NAME}:${IMAGE_VERSION}"
  echo "    4) STOP container ${CONTAINER_NAME}"
  echo "    R) REBUILD/RESTART container ${CONTAINER_NAME} com a imagem ${IMAGE_NAME}:${IMAGE_VERSION}"
  echo "    L) SHOW logs container ${CONTAINER_NAME} rodando"
  echo "    X) SHOW container ${CONTAINER_NAME} rodando"
  echo "    E) EXEC -it no container ${CONTAINER_NAME}"
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
    s|S)
      echo "Saindo... "
      exit 1 ;;
    *) echo "Opcao invalida!!! <ENTER> para sair"
        msgFim ;;
  esac
done
