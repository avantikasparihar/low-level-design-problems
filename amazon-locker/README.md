### Problem Statement:
Design a locker system like Amazon Locker where delivery drivers can deposit packages and customers can pick them up using a code.

### Requirements:

1. Delivery partner can deposit using order-id and partner-id.
2. Customer can pick by using a code shown to them in the order details.
3. Customer needs locker details(location, locker-id, pw) to access the locker. 
4. Concurrency handling is not needed for the operations as it's physically not possible for two people to access the locker at the same time.
5. Authentication of delivery partner or the customer is not required.
6. Capacity management for a given location. The inventory has small, medium, large size lockers. Match the parcel size exactly, if the size is not available, reject the deposit request.
