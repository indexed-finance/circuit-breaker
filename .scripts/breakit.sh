#! /bin/bash

# used to stress test the circuit breaker and try to trigger issues


while true; do

    ./circuit-breaker \
        --config.path circuit-breaker-client.yaml \
        contracts \
        pool-swap \
        --token.out 0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC \
        --token.in 0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66 \
        --token.in_amount "20.0" \
        --token.out_amount "1.0" &

    sleep 1.75s

    ./circuit-breaker \
        --config.path circuit-breaker-client.yaml \
        contracts \
        pool-swap \
        --token.out 0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66 \
        --token.in 0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC \
        --token.in_amount "20.0" \
        --token.out_amount "1.0" &

    sleep 1.75s

    ./circuit-breaker \
        --config.path circuit-breaker-client.yaml \
        contracts \
        pool-swap \
        --token.out 0x1EBed07D5ff3d37A0E03C6EAaC57e562e29f0eee \
        --token.in 0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC \
        --token.in_amount "20.0" \
        --token.out_amount "1.0" &

    sleep 1.75s

    ./circuit-breaker \
        --config.path circuit-breaker-client.yaml \
        contracts \
        pool-swap \
        --token.out 0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC \
        --token.in 0x1EBed07D5ff3d37A0E03C6EAaC57e562e29f0eee \
        --token.in_amount "20.0" \
        --token.out_amount "1.0" &

    sleep 1.75s

   ./circuit-breaker \
        --config.path circuit-breaker-client.yaml \
        contracts \
        pool-swap \
        --token.out 0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66 \
        --token.in 0x1EBed07D5ff3d37A0E03C6EAaC57e562e29f0eee \
        --token.in_amount "20.0" \
        --token.out_amount "1.0" &

    sleep 1.75s

   ./circuit-breaker \
        --config.path circuit-breaker-client.yaml \
        contracts \
        pool-swap \
        --token.out 0x1EBed07D5ff3d37A0E03C6EAaC57e562e29f0eee \
        --token.in 0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66 \
        --token.in_amount "20.0" \
        --token.out_amount "1.0" &

    sleep 1.75s

done


# 0x1EBed07D5ff3d37A0E03C6EAaC57e562e29f0eee
# 0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66
# 0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC