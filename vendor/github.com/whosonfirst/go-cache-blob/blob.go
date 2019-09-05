package blob

import (
	"bufio"
	"bytes"
	"context"
	"github.com/aaronland/gocloud-blob-bucket"	
	"github.com/whosonfirst/go-cache"
	"gocloud.dev/blob"
	"io"
	"io/ioutil"
	"sync/atomic"
)

type BlobCache struct {
	cache.Cache
	TTL int64
	bucket    *blob.Bucket
	hits      int64
	misses    int64
	sets      int64
	evictions int64
}

func NewBlobCacheWithDSN(bucket_dsn string) (cache.Cache, error) {

	ctx := context.Background()

	bucket, err := bucket.OpenBucket(ctx, bucket_dsn)

	if err != nil {
		return nil, err
	}

	return NewBlobCacheWithBucket(bucket)
}

func NewBlobCacheWithBucket(bucket *blob.Bucket) (cache.Cache, error) {

	c := BlobCache{
		TTL:            0,		
		bucket: bucket,
		hits:   0,
		misses: 0,
		sets:   0,
		evictions: 0,
	}

	return &c, nil
}

func (c *BlobCache) Name() string {
	return "blob"
}

func (c *BlobCache) Get(ctx context.Context, key string) (io.ReadCloser, error) {

	fh, err := c.bucket.NewReader(ctx, key, nil)

	if err != nil {
		atomic.AddInt64(&c.misses, 1)
		return nil, err
	}

	atomic.AddInt64(&c.hits, 1)
	return fh, nil
}

func (c *BlobCache) Set(ctx context.Context, key string, fh io.ReadCloser) (io.ReadCloser, error) {

	bucket_wr, err := c.bucket.NewWriter(ctx, key, nil)

	if err != nil {
		return nil, err
	}

	// this is not awesome but until we update all the things (and
	// in particular all the go-whosonfirst-readwrite stuff) to be
	// ReadSeekCloser thingies it's what necessary...
	// (20180617/thisisaaronland)

	var b bytes.Buffer
	wr := bufio.NewWriter(&b)

	io.Copy(wr, fh)
	wr.Flush()

	r := bytes.NewReader(b.Bytes())

	_, err = io.Copy(bucket_wr, r)

	if err != nil {
		return nil, err
	}

	err = bucket_wr.Close()

	if err != nil {
		return nil, err
	}

	atomic.AddInt64(&c.sets, 1)

	r.Reset(b.Bytes())
	return ioutil.NopCloser(r), nil
}

func (c *BlobCache) Unset(ctx context.Context, key string) error {

	err := c.bucket.Delete(ctx, key)

	if err != nil {
		return err
	}

	atomic.AddInt64(&c.evictions, 1)
	return nil
}

func (c *BlobCache) Hits() int64 {
	return atomic.LoadInt64(&c.hits)
}

func (c *BlobCache) Misses() int64 {
	return atomic.LoadInt64(&c.misses)
}

func (c *BlobCache) Evictions() int64 {
	return atomic.LoadInt64(&c.evictions)
}

func (c *BlobCache) Size() int64 {

	return c.SizeWithContext(context.Background())
}

func (c *BlobCache) SizeWithContext(ctx context.Context) int64 {

	size := int64(0)

	iter := c.bucket.List(nil)

	for {

		select {
		case <- ctx.Done():
			return -1
		default:
			//
		}
		
		obj, err := iter.Next(ctx)

		if err == io.EOF {
			break
		}

		if err != nil {
			return -1
		}

		size += obj.Size
	}

	return size
}
