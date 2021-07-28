# goshopify

This library is currently far from finished. We will extend it when we need more endpoints for our projects. Feel free to add more endpoints as well.

# Install

```console
go get github.com/jjideenschmiede/goshopify
```

# How to use?

## Add product

With this function, you can add both simple and variant products. Here you can find an example of how to add a simple product. If it is to be a variant product, then variants are also added in the []AddProductBodyVariants slice, but instead of using the Title option, another one such as size or color is used and the variants are assigned one of the values present in the option.

You can find the description from Shopify [here](https://shopify.dev/api/admin/rest/reference/products/product).

```go
// Define request
r := &Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Define body
body := &AddProductBody{
    AddProductBodyProduct{
        Title:       "",
        BodyHtml:    "",
        Vendor:      "",
        ProductType: "",
        Status:      "",
        Tags:        "",
        Images:      []AddProductBodyImages{},
        Variants:    []AddProductBodyVariants{},
        Options:     []AddProductBodyOptions{},
    },
}

// Add new images
body.Product.Images = append(body.Product.Images, AddProductBodyImages{Src: ""})

// Add new variants
body.Product.Variants = append(body.Product.Variants, AddProductBodyVariants{
    Title:               "",
    Price:               "",
    Sku:                 "",
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
body.Product.Options = append(body.Product.Options, AddProductBodyOptions{
    Name:   "Title",
    Values: []string{"Default Title"},
})

// Add new product
product, err := AddProduct(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(product)
}
```

# Help

For help or questions, please contact us directly [here](mailto:info@jj-development.de).