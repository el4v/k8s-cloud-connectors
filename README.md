# Описание процесса

## Установка на свой кластер

*Для установки __Yandex Cloud Connectors__ необходим настроенный кластер [__k8s__](https://kubernetes.io/) 
и установленный [__helm__](https://helm.sh/).*

Проверим работоспособность коннекторов на Container Registry. Для начала нам надо дать права на работу с Container Registry
сервисному аккаунту, под управлением которого работают ноды в нашем кластере.

Это можно сделать через UI:

```Кластер -> Обзор -> Сервисный аккаунт для узлов -> Редактировать аккаунт -> Роли в каталоге -> '+' -> 'container-registry.admin'```

Того же самого результата можно добиться из консоли:

```shell
folder_id=b1g7jvgmf06eel94s22d # Подставить свой
nodegroup_id=catnmbc81hdag58trmgi # Подставить свой

instance_group_id=$(yc managed-kubernetes node-group get --id ${nodegroup_id} --format json | jq -r ".instance_group_id")
service_account_id=$(yc compute instance-group get --id ${instance_group_id} --format json | jq -r ".instance_template.service_account_id")
yc resource-manager folder add-access-binding --id $folder_id --role container-registry.admin --service-account-id $service_account_id
```

Теперь у нод кластера есть права администрировать Container Registry в облаке. Теперь установим контроллер в кластер, добавляем в кластер все необходимые сущности и контроллер:

```shell
make install
```

Теперь попробуем создать какой-нибудь облачный ресурс, например, **Yandex Container Registry**:

```shell
folder_id=$folder_id envsubst < ./connectors/ycr/examples/test-registry.yaml.tmpl | kubectl apply -f -
```

Можно зайти в UI облака и посмотреть, что там создался новый Container Registry. Аналогичную проверку можно выполнить
с помощью консольной команды `yc container registry list`. Теперь удалим Registry и деинсталлируем из кластера 
контроллер:

```shell
folder_id=$folder_id envsubst < ./connectors/ycr/examples/test-registry.yaml.tmpl | kubectl delete -f -
```

Повторно сходив в веб-интерфейс или исполнив команду `yc container registry list` можно увидеть, что Registry
удалён.

Чтобы удалить связку коннекторов из кластера, достаточно выполнить команду:

```shell
make uninstall
```