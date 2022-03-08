# goshopify

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/goshopify.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/goshopify/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/goshopify/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/goshopify)](https://goreportcard.com/report/github.com/jjideenschmiede/goshopify) [![Go Reference](https://pkg.go.dev/badge/github.com/jjideenschmiede/goshopify.svg)](https://pkg.go.dev/github.com/jjideenschmiede/goshopify) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/goshopify) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/) 

This library is currently far from finished. We will extend it when we need more endpoints for our projects. Feel free to add more endpoints as well.

# Install

```console
go get github.com/jjideenschmiede/goshopify
```

# How to use?

## Get a single product

If you want to read out an individual product directly via the Id, you can do this using this function. You can find the description from Shopify [here](https://shopify.dev/api/admin-rest/2021-07/resources/product#[get]/admin/api/2021-07/products/{product_id}.json).

```go
// Define request
r := goshopify.Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Get a product by id
product, err := goshopify.Product(6147990126757, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(product)
}
```

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
        Id:                  6917353078,
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

## Update a variant

If you want to renew an existing variant, this works similarly to creating a variant. **But you use the id of the variant and not the id of the main product.**

You can find the description from Shopify [here](https://shopify.dev/api/admin-rest/2021-07/resources/product-variant#[put]/admin/api/2021-07/variants/{variant_id}.json).

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
        Id:                  6917353078968,
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
variant, err := goshopify.AddProductVariant(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(variant)
}
```

## Get all inventory locations

If you want to read out all inventory locations, you can do this with the following function. You can find the description from Shopify [here](https://shopify.dev/api/admin-rest/2021-07/resources/location#[get]/admin/api/2021-07/locations.json).

```go
// Define request
r := goshopify.Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Get all inventory locations
locations, err := goshopify.InventoryLocations(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(locations)
}
```

## Set an inventory level

If you want to customize a pass, you can do so as follows. For this you need the Inventory Id and the Location Id of the product. You can find the description from Shopify [here](https://shopify.dev/api/admin-rest/2021-07/resources/inventorylevel#[post]/admin/api/2021-07/inventory_levels/set.json).

```go
// Define request
r := goshopify.Request{
    ApiKey:      "",
    ApiPassword: "",
    StoreName:   "",
}

// Define body
body := goshopify.InventoryLevelBody{
    LocationId:      62413209784,
    InventoryItemId: 42744167694520,
    Available:       24,
}

// Get all inventory locations
inventory, err := goshopify.InventoryLevels(body, r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(inventory)
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

## Add fulfillment to order

If the item has been shipped, the following function can be used to submit the tracking number, shipping provider and tracking link.

You can find the description from Shopify [here](https://shopify.dev/api/admin-rest/2022-01/resources/fulfillment#post-orders-order-id-fulfillments).

```go
// Define fulfillment body
body := FulfillmentBody{
    Fulfillment: FulfillmentBodyFulfillment{
        LocationId:      62413209784,
        TrackingNumber:  "123456789010",
        TrackingCompany: "fed ex",
        TrackingUrl:     "https://test.de/tracking/123456789010",
    },
}

// Add new fulfillment
fulfillment, err := AddFulfillment(4344802345144, body, r)
if err != nil {
    log.Fatalln(err)
} else {
    log.Println(fulfillment)
}
```

# Help

For help or questions, please contact us directly [here](mailto:info@jj-development.de).