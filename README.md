# goshopify

This library is currently far from finished. We will extend it when we need more endpoints for our projects. Feel free to add more endpoints as well.

# Install

```console
go get github.com/jjideenschmiede/goshopify
```

# How to use?

## Add a product

With this function, you can add both simple and variant products. Here you can find an example of how to add a simple product. If it is to be a variant product, then variants are also added in the []ProductBodyVariants slice, but instead of using the Title option, another one such as size or color is used and the variants are assigned one of the values present in the option. **It is important that the Id in the ProductBodyProduct struct is 0.**

You can find the description from Shopify [here](https://shopify.dev/api/admin/rest/reference/products/product).

```go
// Define request
r := goshopify.Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Define body
body := goshopify.ProductBody{
    goshopify.ProductBodyProduct{
    	Id:          0,
        Title:       "",
        BodyHtml:    "",
        Vendor:      "",
        ProductType: "",
        Status:      "",
        Tags:        "",
        Images:      []goshopify.ProductBodyImages{},
        Variants:    []goshopify.ProductBodyVariants{},
        Options:     []goshopify.ProductBodyOptions{},
    },
}

// Add new images
body.Product.Images = append(body.Product.Images, goshopify.ProductBodyImages{Src: ""})

// Add new variants
body.Product.Variants = append(body.Product.Variants, goshopify.ProductBodyVariants{
    Title:               "",
    Price:               "",
    Sku:                 "",
    CompareAtPrice       "",
    FulfillmentService:  "",
    InventoryManagement: "",
    Option1:             "Default Title",
    Option2:             "",
    Option3:             "",
    Taxable:             false,
    Barcode:             "",
    Grams:               0,
    Weight:              0,
    WeightUnit:          "",
    InventoryQuantity:   0,
    RequiresShipping:    false,
})

// Add new options
body.Product.Options = append(body.Product.Options, goshopify.ProductBodyOptions{
    Name:   "Title",
    Values: []string{"Default Title"},
})

// Add new product
product, err := goshopify.AddProduct(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(product)
}
```

## Update a product

Updating a product is actually identical to creating a product. The important thing here is that the ID in the ProductBodyProduct struct is the ID of the product. So the whole product can be updated.

The only difference is that a different function must be called.

You can find the description from Shopify [here](https://shopify.dev/api/admin/rest/reference/products/product).

```go
// Add new product
updateProduct, err := goshopify.UpdateProduct(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(product)
}
```

## Delete a product

If you want to remove a product, you can do this using the following function.

The ID of the product must also be specified. Please specify the request struct.

You can find the description from Shopify [here](https://shopify.dev/api/admin/rest/reference/products/product#destroy-2021-07).

```go
// Delete product
err := DeleteProduct(6881118224568, r)
if err != nil {
	fmt.Println(err)
}
```

## Get a list of all product variants

If you want to read out all variants of a product, you can do this as follows. Only the ID of the main product is required for this.

You can find the description from Shopify [here](https://shopify.dev/api/admin-rest/2021-07/resources/product-variant#[get]/admin/api/2021-07/products/{product_id}/variants.json).

```go
// Define request
r := goshopify.Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Get all product variants
productVariants, err := goshopify.ProductVariants(6917353078968, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(productVariants)
}
```

## Create a variant to a product

If you want to create a new variant, you can do so using the following function. The ID of the main product is required.

You can find the description from Shopify [here](https://shopify.dev/api/admin-rest/2021-07/resources/product-variant#[post]/admin/api/2021-07/products/{product_id}/variants.json).

```go
// Define request
r := goshopify.Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Define body
body := ProductVariantBody{
    ProductVariantBodyVariant{
        Title:               "Testartikel",
        Price:               "24,99",
        Sku:                 "21481462121",
        CompareAtPrice:      "48,99",
        FulfillmentService:  "manual",
        InventoryManagement: "shopify",
        Option1:             "42",
        Option2:             "",
        Option3:             "",
        Taxable:             true,
        Barcode:             "",
        Grams:               0,
        Weight:              0,
        WeightUnit:          "",
        InventoryQuantity:   0,
        RequiresShipping:    true,
    },
}

// Create a product variant
variant, err := goshopify.AddProductVariant(6917353078968, body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(variant)
}
```

## Get list of orders since id

To get a list of orders, you can call the following function. The list shows all orders by a certain ID, if you leave this ID at 0, then the orders are displayed from the beginning. 200 orders are always read out at once.

You can find the description from Shopify [here](https://shopify.dev/api/admin/rest/reference/orders/order).

```go
// Get order list after id (200 items)
orders, err := goshopify.Orders(0, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(orders)
}
```

# Help

For help or questions, please contact us directly [here](mailto:info@jj-development.de).