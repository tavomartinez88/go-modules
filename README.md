# go-modules

##Users
module for create users as admin, client or provider.

###Install
go get github.com/tavomartinez88/go-modules/users

####Use
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