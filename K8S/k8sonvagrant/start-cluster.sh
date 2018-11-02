for node in master worker1 worker2;
  do
    echo "starting" $node " ..."
    vagrant up $node
  done 
echo "cluster started."
