resource "aws_dynamodb_table" "connection" {
  name             = replace("${var.system_name_prefix}_connection", "_", "-")
  billing_mode     = "PAY_PER_REQUEST"
  hash_key         = var.dynamodb_hash_key  // パーティションキー
  range_key        = var.dynamodb_range_key // ソートキー
  stream_enabled   = true                   // dynamoDB streamの有効化
  stream_view_type = "NEW_AND_OLD_IMAGES"   //テーブルに値がきたときのclientへの流し方
  attribute {
    name = var.dynamodb_hash_key
    type = "S"
  }

  attribute {
    name = var.dynamodb_range_key
    type = "S"
  }

  attribute {
    name = var.dynamodb_attr_key1
    type = "S"
  }

  ttl {
    attribute_name = var.dynamodb_ttl_key
    enabled        = true
  }

  // 他のキーで検索したいならこれの設定が必要
  global_secondary_index {
    name               = var.dynamodb_gsi_name
    hash_key           = var.dynamodb_attr_key1
    write_capacity     = 3
    read_capacity      = 5
    projection_type    = "INCLUDE"
    non_key_attributes = [var.dynamodb_hash_key, var.dynamodb_range_key]
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    Name        = "${var.system_name_prefix}_dynamodb"
    Environment = var.aws_env
  }
}