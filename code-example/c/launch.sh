
case $1 in
    "1")
        gcc -o build/mainV1 mainV1.c && build/mainV1
    ;;
    "2")
        gcc -o build/mainV2 mainV2.c && build/mainV2
    ;;
    "3")
        gcc -o build/mainV3 mainV3.c && build/mainV3
    ;;
    *)
        echo -e "No script with that number.\n"
    ;;
esac
