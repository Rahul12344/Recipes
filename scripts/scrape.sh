#!/bin/bash


GOOGLE_CLOUD_RECIPES_BUCKET=recipez_recipes
GOOGLE_CLOUD_FOOD_BUCKET=recipez_food

UNPACK_AND_PORT=./scripts/scrapers/port.py

RECIPES_DATASET=https://storage.googleapis.com/recipe-box/recipes_raw.zip
food_dataset=https://storage.googleapis.com/food_dataset.csv

download () {
    curl -X POST --data-binary @[OBJECT_LOCATION] \
        -H "Authorization: Bearer $OAUTH2_TOKEN" \
        -H "Content-Type: $OBJECT_CONTENT_TYPE" \
        "https://storage.googleapis.com/upload/storage/v1/b/$GOOGLE_CLOUD_RECIPES_BUCKET/o?uploadType=media&name=$OBJECT_NAME"
}

port_to_postgres () {
    echo "Setting up postgres table for recipes and porting recipes to database..."
    `$UNPACK_AND_PORT`
}

download && port_to_postgres
