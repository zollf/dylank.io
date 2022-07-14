# Used for ecr deployment

REPO=$1
AWS=$2
REGION=$3
DOCKERFILE=$4
CONTEXT=$5

if [ -z "$REPO" ]; then
  echo "Please enter repository variable as first paramater"
  exit
fi

if [ -z "$AWS" ]; then
  echo "Please enter aws url as second paramater"
  exit
fi

if [ -z "$REGION" ]; then
  $REGION="ap-southeast-2"
fi

if [ -z "$DOCKERFILE" ]; then
  $DOCKERFILE="./Dockerfile"
fi

if [ -z "$CONTEXT" ]; then
  $CONTEXT="."
fi

docker build -f $DOCKERFILE -t $REPO $CONTEXT || exit
docker tag $REPO:latest $AWS/$REPO:latest || exit
docker login -u AWS -p $(aws ecr get-login-password --region $REGION) $AWS || exit
docker push $AWS/$REPO:latest || exit

# Getting all images that are untagged and deleting them
IMAGES_TO_DELETE=$(aws ecr list-images --region $REGION --repository-name $REPO --filter tagStatus=UNTAGGED --query "imageIds[*]" --output json)
if [ "$IMAGES_TO_DELETE" != "[]" ]
  then
    aws ecr batch-delete-image --region $REGION --repository-name $REPO --image-ids "$IMAGES_TO_DELETE" || true
else
    echo "Skipping batch delete, no images to delete."
fi
