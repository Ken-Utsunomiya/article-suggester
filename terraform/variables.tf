variable "aws_region" {
  default = "ap-northeast-1"
}

variable "vpc_cidr_block" {
  default = "10.0.0.0/16"
}

variable "eks_cluster_name" {
  default = "article-suggester-cluster"
}

variable "pod_execution_role_arn" {
  default = "arn:aws:iam::953201013282:user/article-suggester"
}
