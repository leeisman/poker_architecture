
declare -a files=(
  "Kafka_Overview.md"
  "Topic_and_Partition.md"
  "Producer_Config.md"
  "Consumer_Group.md"
  "Kafka_idempotence.md"
  "Kafka_Reliability.md"
  "Replication_and_Durability.md"
  "Failover_and_Recovery.md"
  "Kafka_Broker.md"
  "Kafka_Stream_and_Connect.md"
  "Kafka_vs_Redis_Comparison.md"
)

i=1
for f in "${files[@]}"; do
  if [ -f "$f" ]; then
    new_name=$(printf "%02d_%s" "$i" "$f")
    mv "$f" "$new_name"
    echo "Renamed: $f → $new_name"
  else
    echo "❗ File not found: $f"
  fi
  ((i++))
done