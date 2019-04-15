package azblob

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/xml"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// blockBlobClient is the client for the BlockBlob methods of the Azblob service.
type blockBlobClient struct {
	managementClient
}

// newBlockBlobClient creates an instance of the blockBlobClient client.
func newBlockBlobClient(url url.URL, p pipeline.Pipeline) blockBlobClient {
	return blockBlobClient{newManagementClient(url, p)}
}

// CommitBlockList the Commit Block List operation writes a blob by specifying the list of block IDs that make up the
// blob. In order to be written as part of a blob, a block must have been successfully written to the server in a prior
// Put Block operation. You can call Put Block List to update a blob by uploading only those blocks that have changed,
// then committing the new and existing blocks together. You can do this by specifying whether to commit a block from
// the committed block list or from the uncommitted block list, or to commit the most recently uploaded version of the
// block, whichever list it may belong to.
//
// timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> blobCacheControl is optional. Sets the blob's cache control. If specified,
// this property is stored with the blob and returned with a read request. blobContentType is optional. Sets the blob's
// content type. If specified, this property is stored with the blob and returned with a read request.
// blobContentEncoding is optional. Sets the blob's content encoding. If specified, this property is stored with the
// blob and returned with a read request. blobContentLanguage is optional. Set the blob's content language. If
// specified, this property is stored with the blob and returned with a read request. blobContentMD5 is optional. An
// MD5 hash of the blob content. Note that this hash is not validated, as the hashes for the individual blocks were
// validated when each was uploaded. metadata is optional. Specifies a user-defined name-value pair associated with the
// blob. If no name-value pairs are specified, the operation will copy the metadata from the source blob or file to the
// destination blob. If one or more name-value pairs are specified, the destination blob is created with the specified
// metadata, and metadata is not copied from the source blob or file. Note that beginning with version 2009-09-19,
// metadata names must adhere to the naming rules for C# identifiers. See Naming and Referencing Containers, Blobs, and
// Metadata for more information. leaseID is if specified, the operation only succeeds if the container's lease is
// active and matches this ID. blobContentDisposition is optional. Sets the blob's Content-Disposition header.
// ifModifiedSince is specify this header value to operate only on a blob if it has been modified since the specified
// date/time. ifUnmodifiedSince is specify this header value to operate only on a blob if it has not been modified
// since the specified date/time. ifMatches is specify an ETag value to operate only on blobs with a matching value.
// ifNoneMatch is specify an ETag value to operate only on blobs without a matching value. requestID is provides a
// client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when storage
// analytics logging is enabled.
func (client blockBlobClient) CommitBlockList(ctx context.Context, blocks BlockLookupList, timeout *int32, blobCacheControl *string, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 []byte, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (*BlockBlobCommitBlockListResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}},
		{targetValue: metadata,
			constraints: []constraint{{target: "metadata", name: null, rule: false,
				chain: []constraint{{target: "metadata", name: pattern, rule: `^[a-zA-Z]+$`, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.commitBlockListPreparer(blocks, timeout, blobCacheControl, blobContentType, blobContentEncoding, blobContentLanguage, blobContentMD5, metadata, leaseID, blobContentDisposition, ifModifiedSince, ifUnmodifiedSince, ifMatches, ifNoneMatch, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.commitBlockListResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockBlobCommitBlockListResponse), err
}

// commitBlockListPreparer prepares the CommitBlockList request.
func (client blockBlobClient) commitBlockListPreparer(blocks BlockLookupList, timeout *int32, blobCacheControl *string, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 []byte, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "blocklist")
	req.URL.RawQuery = params.Encode()
	if blobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobCacheControl)
	}
	if blobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobContentType)
	}
	if blobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobContentEncoding)
	}
	if blobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobContentLanguage)
	}
	if blobContentMD5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(blobContentMD5))
	}
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	if blobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobContentDisposition)
	}
	if ifModifiedSince != nil {
		req.Header.Set("If-Modified-Since", (*ifModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", (*ifUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifMatches != nil {
		req.Header.Set("If-Match", string(*ifMatches))
	}
	if ifNoneMatch != nil {
		req.Header.Set("If-None-Match", string(*ifNoneMatch))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	b, err := xml.Marshal(blocks)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/xml")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// commitBlockListResponder handles the response to the CommitBlockList request.
func (client blockBlobClient) commitBlockListResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &BlockBlobCommitBlockListResponse{rawResponse: resp.Response()}, err
}

// GetBlockList the Get Block List operation retrieves the list of blocks that have been uploaded as part of a block
// blob
//
// listType is specifies whether to return the list of committed blocks, the list of uncommitted blocks, or both lists
// together. snapshot is the snapshot parameter is an opaque DateTime value that, when present, specifies the blob
// snapshot to retrieve. For more information on working with blob snapshots, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/creating-a-snapshot-of-a-blob">Creating
// a Snapshot of a Blob.</a> timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> leaseID is if specified, the operation only succeeds if the container's
// lease is active and matches this ID. requestID is provides a client-generated, opaque value with a 1 KB character
// limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client blockBlobClient) GetBlockList(ctx context.Context, listType BlockListType, snapshot *string, timeout *int32, leaseID *string, requestID *string) (*BlockList, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.getBlockListPreparer(listType, snapshot, timeout, leaseID, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getBlockListResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockList), err
}

// getBlockListPreparer prepares the GetBlockList request.
func (client blockBlobClient) getBlockListPreparer(listType BlockListType, snapshot *string, timeout *int32, leaseID *string, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("GET", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if snapshot != nil && len(*snapshot) > 0 {
		params.Set("snapshot", *snapshot)
	}
	params.Set("blocklisttype", string(listType))
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "blocklist")
	req.URL.RawQuery = params.Encode()
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// getBlockListResponder handles the response to the GetBlockList request.
func (client blockBlobClient) getBlockListResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &BlockList{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		b = removeBOM(b)
		err = xml.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// StageBlock the Stage Block operation creates a new block to be committed as part of a blob
//
// blockID is a valid Base64 string value that identifies the block. Prior to encoding, the string must be less than or
// equal to 64 bytes in size. For a given blob, the length of the value specified for the blockid parameter must be the
// same size for each block. contentLength is the length of the request. body is initial data body will be closed upon
// successful return. Callers should ensure closure when receiving an error.timeout is the timeout parameter is
// expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> leaseID is if specified, the operation only succeeds if the container's
// lease is active and matches this ID. requestID is provides a client-generated, opaque value with a 1 KB character
// limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client blockBlobClient) StageBlock(ctx context.Context, blockID string, contentLength int64, body io.ReadSeeker, timeout *int32, leaseID *string, requestID *string) (*BlockBlobStageBlockResponse, error) {
	if err := validate([]validation{
		{targetValue: body,
			constraints: []constraint{{target: "body", name: null, rule: true, chain: nil}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.stageBlockPreparer(blockID, contentLength, body, timeout, leaseID, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.stageBlockResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockBlobStageBlockResponse), err
}

// stageBlockPreparer prepares the StageBlock request.
func (client blockBlobClient) stageBlockPreparer(blockID string, contentLength int64, body io.ReadSeeker, timeout *int32, leaseID *string, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("blockid", blockID)
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "block")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// stageBlockResponder handles the response to the StageBlock request.
func (client blockBlobClient) stageBlockResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &BlockBlobStageBlockResponse{rawResponse: resp.Response()}, err
}

// StageBlockFromURL the Stage Block operation creates a new block to be committed as part of a blob where the contents
// are read from a URL.
//
// blockID is a valid Base64 string value that identifies the block. Prior to encoding, the string must be less than or
// equal to 64 bytes in size. For a given blob, the length of the value specified for the blockid parameter must be the
// same size for each block. contentLength is the length of the request. sourceURL is specifiy an URL to the copy
// source. sourceRange is bytes of source data in the specified range. sourceContentMD5 is specify the md5 calculated
// for the range of bytes that must be read from the copy source. timeout is the timeout parameter is expressed in
// seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> leaseID is if specified, the operation only succeeds if the container's
// lease is active and matches this ID. requestID is provides a client-generated, opaque value with a 1 KB character
// limit that is recorded in the analytics logs when storage analytics logging is enabled.
func (client blockBlobClient) StageBlockFromURL(ctx context.Context, blockID string, contentLength int64, sourceURL *string, sourceRange *string, sourceContentMD5 []byte, timeout *int32, leaseID *string, requestID *string) (*BlockBlobStageBlockFromURLResponse, error) {
	if err := validate([]validation{
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.stageBlockFromURLPreparer(blockID, contentLength, sourceURL, sourceRange, sourceContentMD5, timeout, leaseID, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.stageBlockFromURLResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockBlobStageBlockFromURLResponse), err
}

// stageBlockFromURLPreparer prepares the StageBlockFromURL request.
func (client blockBlobClient) stageBlockFromURLPreparer(blockID string, contentLength int64, sourceURL *string, sourceRange *string, sourceContentMD5 []byte, timeout *int32, leaseID *string, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("blockid", blockID)
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	params.Set("comp", "block")
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if sourceURL != nil {
		req.Header.Set("x-ms-copy-source", *sourceURL)
	}
	if sourceRange != nil {
		req.Header.Set("x-ms-source-range", *sourceRange)
	}
	if sourceContentMD5 != nil {
		req.Header.Set("x-ms-source-content-md5", base64.StdEncoding.EncodeToString(sourceContentMD5))
	}
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// stageBlockFromURLResponder handles the response to the StageBlockFromURL request.
func (client blockBlobClient) stageBlockFromURLResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &BlockBlobStageBlockFromURLResponse{rawResponse: resp.Response()}, err
}

// Upload the Upload Block Blob operation updates the content of an existing block blob. Updating an existing block
// blob overwrites any existing metadata on the blob. Partial updates are not supported with Put Blob; the content of
// the existing blob is overwritten with the content of the new blob. To perform a partial update of the content of a
// block blob, use the Put Block List operation.
//
// body is initial data body will be closed upon successful return. Callers should ensure closure when receiving an
// error.contentLength is the length of the request. timeout is the timeout parameter is expressed in seconds. For more
// information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> blobContentType is optional. Sets the blob's content type. If specified,
// this property is stored with the blob and returned with a read request. blobContentEncoding is optional. Sets the
// blob's content encoding. If specified, this property is stored with the blob and returned with a read request.
// blobContentLanguage is optional. Set the blob's content language. If specified, this property is stored with the
// blob and returned with a read request. blobContentMD5 is optional. An MD5 hash of the blob content. Note that this
// hash is not validated, as the hashes for the individual blocks were validated when each was uploaded.
// blobCacheControl is optional. Sets the blob's cache control. If specified, this property is stored with the blob and
// returned with a read request. metadata is optional. Specifies a user-defined name-value pair associated with the
// blob. If no name-value pairs are specified, the operation will copy the metadata from the source blob or file to the
// destination blob. If one or more name-value pairs are specified, the destination blob is created with the specified
// metadata, and metadata is not copied from the source blob or file. Note that beginning with version 2009-09-19,
// metadata names must adhere to the naming rules for C# identifiers. See Naming and Referencing Containers, Blobs, and
// Metadata for more information. leaseID is if specified, the operation only succeeds if the container's lease is
// active and matches this ID. blobContentDisposition is optional. Sets the blob's Content-Disposition header.
// ifModifiedSince is specify this header value to operate only on a blob if it has been modified since the specified
// date/time. ifUnmodifiedSince is specify this header value to operate only on a blob if it has not been modified
// since the specified date/time. ifMatches is specify an ETag value to operate only on blobs with a matching value.
// ifNoneMatch is specify an ETag value to operate only on blobs without a matching value. requestID is provides a
// client-generated, opaque value with a 1 KB character limit that is recorded in the analytics logs when storage
// analytics logging is enabled.
func (client blockBlobClient) Upload(ctx context.Context, body io.ReadSeeker, contentLength int64, timeout *int32, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 []byte, blobCacheControl *string, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (*BlockBlobUploadResponse, error) {
	if err := validate([]validation{
		{targetValue: body,
			constraints: []constraint{{target: "body", name: null, rule: true, chain: nil}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}},
		{targetValue: metadata,
			constraints: []constraint{{target: "metadata", name: null, rule: false,
				chain: []constraint{{target: "metadata", name: pattern, rule: `^[a-zA-Z]+$`, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.uploadPreparer(body, contentLength, timeout, blobContentType, blobContentEncoding, blobContentLanguage, blobContentMD5, blobCacheControl, metadata, leaseID, blobContentDisposition, ifModifiedSince, ifUnmodifiedSince, ifMatches, ifNoneMatch, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.uploadResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*BlockBlobUploadResponse), err
}

// uploadPreparer prepares the Upload request.
func (client blockBlobClient) uploadPreparer(body io.ReadSeeker, contentLength int64, timeout *int32, blobContentType *string, blobContentEncoding *string, blobContentLanguage *string, blobContentMD5 []byte, blobCacheControl *string, metadata map[string]string, leaseID *string, blobContentDisposition *string, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", strconv.FormatInt(int64(*timeout), 10))
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	if blobContentType != nil {
		req.Header.Set("x-ms-blob-content-type", *blobContentType)
	}
	if blobContentEncoding != nil {
		req.Header.Set("x-ms-blob-content-encoding", *blobContentEncoding)
	}
	if blobContentLanguage != nil {
		req.Header.Set("x-ms-blob-content-language", *blobContentLanguage)
	}
	if blobContentMD5 != nil {
		req.Header.Set("x-ms-blob-content-md5", base64.StdEncoding.EncodeToString(blobContentMD5))
	}
	if blobCacheControl != nil {
		req.Header.Set("x-ms-blob-cache-control", *blobCacheControl)
	}
	if metadata != nil {
		for k, v := range metadata {
			req.Header.Set("x-ms-meta-"+k, v)
		}
	}
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	if blobContentDisposition != nil {
		req.Header.Set("x-ms-blob-content-disposition", *blobContentDisposition)
	}
	if ifModifiedSince != nil {
		req.Header.Set("If-Modified-Since", (*ifModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", (*ifUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifMatches != nil {
		req.Header.Set("If-Match", string(*ifMatches))
	}
	if ifNoneMatch != nil {
		req.Header.Set("If-None-Match", string(*ifNoneMatch))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	req.Header.Set("x-ms-blob-type", "BlockBlob")
	return req, nil
}

// uploadResponder handles the response to the Upload request.
func (client blockBlobClient) uploadResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	io.Copy(ioutil.Discard, resp.Response().Body)
	resp.Response().Body.Close()
	return &BlockBlobUploadResponse{rawResponse: resp.Response()}, err
}
