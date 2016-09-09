package "fmt"

func CopyFile(dstName, srcName string) (written int 64, err error)
{
	src,err:= os.Open(srcName)
	if err!=nil {
	return
	}
	defer src.Close()
	 
	dst, err:=os.Create(dstName)
	 if err!= nil {
	    return
	 }
	 defer dst.Close()
	 
	 written, err =io.Copy(dst,src)
	 
	// dst.Close()
//src.Close()
	  
	 return
}