## 옵저버 패턴
```
    옵저버 패턴 : publisher와 subscriber의 상관관계를 갖고 있는 디자인 패턴
    subscription(구독) 이라는 메커니즘을 더해 publisher class가 구독 / 해지를 할 수 있도록 함
```

모든 구독 객체는 사전에 정의된 인터페이스를 implement 하고, 인터페이스에 정의된 방식으로만 통신해야한다.

publisher 쪽에서는 notification과 관련된 parameter를 정의해야한다. 

1. 옵저버 패턴에서는 한개의 관찰 대상자(Subject)와 여러개의 관찰자(Observer)로 일 대 다 관계로 구성되어있다.
2. 옵저버 패턴에서는 관찰 대상(Subject)의 상태가 바뀌면 변경사항을 옵저버에 통보한다 (notify observer)
3. Subject로부터 통보를 받은 Observer는 값을 바꿀 수도 있고, 삭제 하는 등 적절히 대응한다 (update)
4. Observer들은 언제든 Subject의 그룹에서 추가/삭제 될 수 있다. Subject 그룹에 추가되면 Subject로부터 정보를 전달받게 되며, 그룹에서 삭제될 경우 정보를 받아볼 수 없다. 