# What is Deadlock?

二つのトランザクションが二つ以上の資源の確保をめぐって互いに相手を待つ状態となり、そこから先へ処理が進まなくなることを Deadlock（デッドロック）という。

# Entity Relationship Diagram

![EntityRelationshipDiagram](images/dbdiagram.jpg "EntityRelationshipDiagram")

例えば、account1 から account2 へ $10 送金することを想定する。

作成されるレコードは下記の三つ。

1. 【transfers】
   from_account_id：1、to_account_id：2、amount：10 という値が入ったレコード
2. 【entries】
   account_id：1、amount：-10 という値が入ったレコード
3. 【entries】
   account_id：2、amount：10 という値が入ったレコード
