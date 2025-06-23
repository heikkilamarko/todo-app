variable "azure_subscription_id" {}
variable "name" {}
variable "location" {}
variable "owner" {}
variable "public_key_file_path" {}

provider "azurerm" {
  features {
    log_analytics_workspace {
      permanently_delete_on_destroy = true
    }
  }
  subscription_id = var.azure_subscription_id
}

resource "azurerm_resource_group" "app" {
  name     = "rg-${var.name}"
  location = var.location

  tags = {
    owner = var.owner
  }
}

resource "azurerm_virtual_network" "app" {
  name                = "vnet-${var.name}"
  resource_group_name = azurerm_resource_group.app.name
  location            = azurerm_resource_group.app.location
  address_space       = ["10.0.0.0/16"]

  tags = {
    environment = "demo"
  }
}

resource "azurerm_subnet" "app" {
  name                 = "snet-${var.name}"
  resource_group_name  = azurerm_resource_group.app.name
  virtual_network_name = azurerm_virtual_network.app.name
  address_prefixes     = ["10.0.3.0/24"]
}

resource "azurerm_public_ip" "app" {
  name                = "pip-${var.name}"
  resource_group_name = azurerm_resource_group.app.name
  location            = azurerm_resource_group.app.location
  allocation_method   = "Static"
}

resource "azurerm_network_interface" "app" {
  name                = "nic-${var.name}"
  resource_group_name = azurerm_resource_group.app.name
  location            = azurerm_resource_group.app.location

  ip_configuration {
    name                          = "ipconfig"
    subnet_id                     = azurerm_subnet.app.id
    private_ip_address_allocation = "Dynamic"
    public_ip_address_id          = azurerm_public_ip.app.id
  }
}

resource "azurerm_network_security_group" "app" {
  name                = "nsg-${var.name}"
  resource_group_name = azurerm_resource_group.app.name
  location            = azurerm_resource_group.app.location

  security_rule {
    name                       = "SSH"
    priority                   = 100
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
    priority                   = 110
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "443"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }
}

resource "azurerm_network_interface_security_group_association" "app" {
  network_interface_id      = azurerm_network_interface.app.id
  network_security_group_id = azurerm_network_security_group.app.id
}

resource "azurerm_linux_virtual_machine" "app" {
  name                            = "vm-${var.name}"
  resource_group_name             = azurerm_resource_group.app.name
  location                        = azurerm_resource_group.app.location
  network_interface_ids           = [azurerm_network_interface.app.id]
  size                            = "Standard_D2as_v5"
  admin_username                  = "azureuser"
  disable_password_authentication = true

  admin_ssh_key {
    username   = "azureuser"
    public_key = file("${var.public_key_file_path}")
  }

  os_disk {
    name                 = "disk-${var.name}"
    caching              = "ReadWrite"
    storage_account_type = "Premium_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "ubuntu-24_04-lts"
    sku       = "server"
    version   = "latest"
  }

  lifecycle {
    ignore_changes = [
      admin_ssh_key,
    ]
  }
}

output "vm_public_ip" {
  value = azurerm_public_ip.app.ip_address
}

output "vm_admin_username" {
  value = azurerm_linux_virtual_machine.app.admin_username
}
