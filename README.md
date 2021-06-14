# go-modules

## Users
module for create users as admin, client or provider.

### Install
go get github.com/tavomartinez88/go-modules/users

#### Use
For create an admin :
- create an instance admin builder newAdminBuilder()
- Set Username and Password
- Create an director's instance with builder as parameter
    - newDirector(builder)
- create an user director.buildUser()

For create an client :
- create an instance admin builder newClientBuilder()
- Set Username and Password
- [optional] if you want set an client's address and phone use fields:
    - Address
    - Phone 
- Create an director's instance with builder as parameter
    - newDirector(builder)
- create an user director.buildUser()   

For create an provider :
- create an instance admin builder newProviderBuilder()
- Set Username and Password
- [optional] if you want set an client's address and phone use fields as slices:
    - Addresses
    - Phones 
    - AttentionDays
- Create an director's instance with builder as parameter
    - newDirector(builder)
- create an user director.buildUser()  

#### For more info
- run locally: godoc -http=:6060
- access more info about modules:
    -  [doc](http://localhost:6060/pkg/github.com/tavomartinez88/go-modules/users/)    
    
## Categories and SubCategories
module for create categories and sub-categories

### Install
go get github.com/tavomartinez88/go-modules/categories   

#### Use
- Create builder:
    - For category use the func builders.CreateCategoryBuild()
    - For sub-category use the func builders.CreateSubCategoryBuild()
- Create director:
    - CreateCategoryDirector(builder)
- Generate data using director.Build(name, refId, isAvailable)

#### For more info
- run locally: godoc -http=:6060
- access more info about modules:
    -  [doc](http://localhost:6060/pkg/github.com/tavomartinez88/go-modules/categories/) 
    
## Products
module for create categories and sub-categories

### Install
go get github.com/tavomartinez88/go-modules/products
#### Use
- Create builder:
    - Use the func builders.CreateProductBuilder()
    - Set fields name, price, description short, description large, hasStock, and IdRef(it's reference to category or subcategory)
- Create director:
    - CreateProductDirector(builder)

#### For more info
- run locally: godoc -http=:6060
- access more info about modules:
    -  [doc](http://localhost:6060/pkg/github.com/tavomartinez88/go-modules/products/)
    
              