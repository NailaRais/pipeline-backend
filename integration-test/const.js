export const dogImg = open(`${__ENV.TEST_FOLDER_ABS_PATH}/integration-test/data/dog.jpg`, "b");

export const det_model = open(`${__ENV.TEST_FOLDER_ABS_PATH}/integration-test/data/dummy-det-model.zip`, "b");
export const model_def_name = "model-definitions/github"
export const model_def_uid = "909c3278-f7d1-461c-9352-87741bef11d3"
export const model_id = "dummy-det"
export const model_instance_id = "latest"

export const detSyncRecipe = {
  recipe: {
    source: "source-connectors/source-http",
    model_instances: [`models/${model_id}/instances/${model_instance_id}`],
    destination: "destination-connectors/destination-http"
  },
};

export const dstCSVConnID = "some-cool-name-for-dst-csv-connector"

export const detAsyncRecipe = {
  recipe: {
    source: "source-connectors/source-http",
    model_instances: [`models/${model_id}/instances/${model_instance_id}`],
    destination:`destination-connectors/${dstCSVConnID}`
  },
};
