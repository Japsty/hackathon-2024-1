#!/bin/bash
declare -a users=(a.sulaev vpersiyanova d.fedorova vvjkee314 aleksandrov_aa NervousVoid zizaad fzastahov atlasafri BurtsE booec Katetcpip ell1jah shmalens Arugaf poliorang alexey.zakharenkov Alekseizor imaximus3 IvanKaliuzhny Grandeas samarec1812 Yorlend anton.kravtsov.31 krrrsnv deraswer lptnkv ilya-malyshev-2002 sh4rkizz deal_m BlackGoose mao360 Fizinftopixxx max1mn deeezy ocmaxim97 DimasXT _rauzh_ DmitriyRogov LeGushka sjava19021 bloomou drakonnata113 MrSergeevNikita sgurman ISkalchenkov _sobms_ StephanSok SirPerrier iYroglif blackberryBush alpha1811 rayneraido RCNRC levbara1 Artyom1363 alexwerben gassannov nokrolikno mitgottwirstduherrschen maxxximus_prime makezh kaizer-nurik Adefe v-mk-s p0rtale ilya868 superAIyah SimRoman sibwa0 Totenkaf rkhazh)
declare password="dsuyuyduaudijijhugvydsuyuyduaudijijhugvy"


# scp -i description/vm_access_key.pem description/vm_initusers.sh ubuntu@89.208.199.239:~/vm_initusers.sh
# ssh -i description/vm_access_key.pem ubuntu@89.208.199.239 'sudo ./vm_initusers.sh'

: > newusers.txt
for user in "${users[@]}"
do  
    echo $user:$password::::/home/$user:/bin/bash >> newusers.txt
done
chmod 600 newusers.txt
newusers newusers.txt

for user in "${users[@]}"
do
    :
    usermod -aG sudo $user
    rm -rf /home/$user/.ssh
    rsync --archive --chown=$user:$user /home/ubuntu/.ssh /home/$user
done
