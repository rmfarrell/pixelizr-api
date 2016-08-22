provider "aws" {
  access_key="${var.access_key}"
  secret_key="${var.secret_key}"
  region="${var.region}"
}

resource "aws_s3_bucket" "src-bucket" {
  bucket = "pixelizr"
  acl    = "public-read-write"
}

resource "aws_api_gateway_rest_api" "PixelizrApi" {
  name = "Pixelizer API"
}

resource "aws_api_gateway_resource" "Pixelizr" {
  rest_api_id = "${aws_api_gateway_rest_api.PixelizrApi.id}"
  parent_id = "${aws_api_gateway_rest_api.PixelizrApi.root_resource_id}"
  path_part = "pixelizr"
}

resource "aws_api_gateway_method" "PixelizrCreate" {
  rest_api_id = "${aws_api_gateway_rest_api.PixelizrApi.id}"
  resource_id = "${aws_api_gateway_resource.Pixelizr.id}"
  http_method = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "PixelizrCreateIntegration" {
  rest_api_id = "${aws_api_gateway_rest_api.PixelizrApi.id}"
  resource_id = "${aws_api_gateway_resource.Pixelizr.id}"
  http_method = "${aws_api_gateway_method.PixelizrCreate.http_method}"
  type = "AWS"
  uri = "arn:aws:apigateway:${var.region}:lambda:path/2015-03-31/functions/pixelizer-api_hello/invocation"
  credentials = "${var.role_arn}"
  integration_http_method = "${aws_api_gateway_method.PixelizrCreate.http_method}"
}

resource "aws_api_gateway_deployment" "MyDemoDeployment" {
  depends_on = ["aws_api_gateway_method.PixelizrCreate"]
  rest_api_id = "${aws_api_gateway_rest_api.PixelizrApi.id}"
  stage_name = "prod"
}