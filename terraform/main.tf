terraform {
  cloud {
    organization = "kenu"
    workspaces {
      name = "article-suggester"
    }
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.6"
    }
  }
  required_version = ">= 1.2.0"
}

# Configure the AWS provider
provider "aws" {
  region = var.aws_region
}

# VPC
# ---------------------------------------
resource "aws_vpc" "article_suggester_vpc" {
  cidr_block = var.vpc_cidr_block
}

resource "aws_subnet" "article_suggester_subnet" {
  vpc_id     = aws_vpc.article_suggester_vpc.id
  cidr_block = "10.0.1.0/24"
}

# EKS
# ---------------------------------------
resource "aws_eks_cluster" "article_suggester_cluster" {
  name     = "article-suggester-cluster"
  role_arn = aws_iam_role.eks_cluster.arn
  vpc_config {
    subnet_ids = [aws_subnet.article_suggester_subnet.id]
  }
}

resource "aws_iam_role" "eks_cluster" {
  name = "eks-cluster"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "eks.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "eks_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.eks_cluster.name
}

# Fargate
# ---------------------------------------
resource "aws_eks_fargate_profile" "example_fargate_profile" {
  cluster_name           = aws_eks_cluster.article_suggester_cluster.name
  fargate_profile_name   = "article-suggester-fargate-profile"
  pod_execution_role_arn = var.pod_execution_role_arn
  subnet_ids             = [aws_subnet.article_suggester_subnet.id]
  selector {
    namespace = "default"
  }
}
