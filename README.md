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
r := &goshopify.Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Define body
body := &goshopify.ProductBody{
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

## Get list of products since id

To get a list of orders, you can call the following function. The list shows all orders by a certain ID, if you leave this ID at 0, then the orders are displayed from the beginning. 200 orders are always read out at once.

You can find the description from Shopify [here](https://shopify.dev/api/admin/rest/reference/orders/order).

```go
// Get order list after id (200 items)
orders, err := Orders(0, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(orders)
}
```

# Help

For help or questions, please contact us directly [here](mailto:info@jj-development.de).