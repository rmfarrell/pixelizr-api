variable "access_key" {}
variable "secret_key" {}
variable "role_arn" {}
variable "region" {
  default = "us-east-1"
}

variable "apex_function_hello" {
  type="string"
  default="arn:aws:lambda:xxxx:xxxxxxx:function:xxxx_hello"
}