import os
import subprocess

proto_folder = "proto/"

proto_files = [file for file in os.listdir(proto_folder) if file.endswith(".proto")]

for proto_file in proto_files:
    output_dir = os.path.join("internal/transport/rpc/pb", os.path.splitext(proto_file)[0])
    
    subprocess.run(["npx", "buf", "generate", "--path", os.path.join(proto_folder, proto_file), "--output", output_dir], check=True)
    
    go_files = [file for file in os.listdir(output_dir+"/proto") if file.endswith(".go")]
    for file in go_files:
        subprocess.run(["mv", os.path.join(output_dir, "proto", file), output_dir+"/"])
        
     
    subprocess.run(["mv", os.path.join(output_dir, "api/proto", os.path.splitext(proto_file)[0]+".swagger.json"), "api/"])

    subprocess.run(["rm", "-rf", os.path.join(output_dir, "api")], check=True)
    subprocess.run(["rm", "-rf", os.path.join(output_dir, "proto")], check=True)