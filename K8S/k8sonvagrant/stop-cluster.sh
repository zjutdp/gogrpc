for node in master worker1 worker2;
  do
    echo "stopping " $node " ..."
    vagrant halt $node
  done 
echo "cluster stopped."
