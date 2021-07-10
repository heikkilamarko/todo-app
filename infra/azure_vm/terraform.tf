terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>2.0"
    }
  }
}

provider "azurerm" {
  features {}
}

variable "name" {}
variable "location" {}
variable "domain_name_label" {}
variable "public_key_file_path" {}
variable "private_key_file_path" {}

resource "azurerm_resource_group" "app" {
  name     = "rg-${var.name}"
  location = var.location

  tags = {
    environment = "demo"
  }
}

resource "azurerm_virtual_network" "app" {
  name                = "vnet-${var.name}"
  address_space       = ["10.4.0.0/16"]
  location            = azurerm_resource_group.app.location
  resource_group_name = azurerm_resource_group.app.name

  tags = {
    environment = "demo"
  }
}

resource "azurerm_subnet" "app" {
  name                 = "snet-${var.name}"
  resource_group_name  = azurerm_resource_group.app.name
  virtual_network_name = azurerm_virtual_network.app.name
  address_prefixes     = ["10.4.0.0/24"]
}

resource "azurerm_public_ip" "app" {
  name                = "pip-${var.name}"
  location            = azurerm_resource_group.app.location
  resource_group_name = azurerm_resource_group.app.name
  allocation_method   = "Dynamic"
  domain_name_label   = var.domain_name_label

  tags = {
    environment = "demo"
  }
}

resource "azurerm_network_security_group" "app" {
  name                = "nsg-${var.name}"
  location            = azurerm_resource_group.app.location
  resource_group_name = azurerm_resource_group.app.name

  security_rule {
    name                       = "SSH"
    priority                   = 300
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "22"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "HTTPS"
    priority                   = 310
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "443"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "Keycloak"
    priority                   = 320
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "8443"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  security_rule {
    name                       = "MinIO"
    priority                   = 330
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "9443"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  tags = {
    environment = "demo"
  }
}

resource "azurerm_network_interface" "app" {
  name                = "nic-${var.name}"
  location            = azurerm_resource_group.app.location
  resource_group_name = azurerm_resource_group.app.name

  ip_configuration {
    name                          = "ipconfig"
    subnet_id                     = azurerm_subnet.app.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.app.id
  }

  tags = {
    environment = "demo"
  }
}

resource "azurerm_network_interface_security_group_association" "app" {
  network_interface_id      = azurerm_network_interface.app.id
  network_security_group_id = azurerm_network_security_group.app.id
}

resource "azurerm_storage_account" "app" {
  name                     = "st${replace(var.name, "-", "")}001"
  resource_group_name      = azurerm_resource_group.app.name
  location                 = azurerm_resource_group.app.location
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    environment = "demo"
  }
}

resource "azurerm_linux_virtual_machine" "app" {
  name                            = "vm-${var.name}"
  location                        = azurerm_resource_group.app.location
  resource_group_name             = azurerm_resource_group.app.name
  network_interface_ids           = [azurerm_network_interface.app.id]
  size                            = "Standard_D2s_v3"
  admin_username                  = "azureuser"
  disable_password_authentication = true

  os_disk {
    name                 = "disk-${var.name}"
    caching              = "ReadWrite"
    storage_account_type = "Premium_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-focal"
    sku       = "20_04-lts"
    version   = "latest"
  }

  admin_ssh_key {
    username   = "azureuser"
    public_key = file("${var.public_key_file_path}")
  }

  boot_diagnostics {
    storage_account_uri = azurerm_storage_account.app.primary_blob_endpoint
  }

  provisioner "remote-exec" {
    script = "install_docker.sh"

    connection {
      type        = "ssh"
      user        = "azureuser"
      host        = "${var.domain_name_label}.westeurope.cloudapp.azure.com"
      private_key = file("${var.private_key_file_path}")
    }
  }

  tags = {
    environment = "demo"
  }
}
