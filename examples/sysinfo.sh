#!/bin/bash

# Example: A useful extension that shows system information

echo "💻 System Information"
echo "===================="
echo ""
echo "Hostname: $(hostname)"
echo "OS: $(uname -s)"
echo "Kernel: $(uname -r)"
echo "Architecture: $(uname -m)"
echo ""
echo "Current User: $USER"
echo "Home Directory: $HOME"
echo "Current Directory: $(pwd)"
echo ""
echo "Date/Time: $(date)"
echo ""
echo "Disk Usage:"
df -h / | tail -n 1
echo ""
echo "Memory Usage:"
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    vm_stat | head -n 10
else
    # Linux
    free -h
fi
