// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImagesRequest
type DescribeFpgaImagesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	//
	//    * create-time - The creation time of the AFI.
	//
	//    * fpga-image-id - The FPGA image identifier (AFI ID).
	//
	//    * fpga-image-global-id - The global FPGA image identifier (AGFI ID).
	//
	//    * name - The name of the AFI.
	//
	//    * owner-id - The AWS account ID of the AFI owner.
	//
	//    * product-code - The product code.
	//
	//    * shell-version - The version of the AWS Shell that was used to create
	//    the bitstream.
	//
	//    * state - The state of the AFI (pending | failed | available | unavailable).
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	//
	//    * update-time - The time of the most recent update.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The AFI IDs.
	FpgaImageIds []string `locationName:"FpgaImageId" locationNameList:"item" type:"list"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `min:"1" type:"string"`

	// Filters the AFI by owner. Specify an AWS account ID, self (owner is the sender
	// of the request), or an AWS owner alias (valid values are amazon | aws-marketplace).
	Owners []string `locationName:"Owner" locationNameList:"Owner" type:"list"`
}

// String returns the string representation
func (s DescribeFpgaImagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFpgaImagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFpgaImagesInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImagesResult
type DescribeFpgaImagesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the FPGA images.
	FpgaImages []FpgaImage `locationName:"fpgaImageSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeFpgaImagesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeFpgaImages = "DescribeFpgaImages"

// DescribeFpgaImagesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the Amazon FPGA Images (AFIs) available to you. These include public
// AFIs, private AFIs that you own, and AFIs owned by other AWS accounts for
// which you have load permissions.
//
//    // Example sending a request using DescribeFpgaImagesRequest.
//    req := client.DescribeFpgaImagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImages
func (c *Client) DescribeFpgaImagesRequest(input *DescribeFpgaImagesInput) DescribeFpgaImagesRequest {
	op := &aws.Operation{
		Name:       opDescribeFpgaImages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeFpgaImagesInput{}
	}

	req := c.newRequest(op, input, &DescribeFpgaImagesOutput{})
	return DescribeFpgaImagesRequest{Request: req, Input: input, Copy: c.DescribeFpgaImagesRequest}
}

// DescribeFpgaImagesRequest is the request type for the
// DescribeFpgaImages API operation.
type DescribeFpgaImagesRequest struct {
	*aws.Request
	Input *DescribeFpgaImagesInput
	Copy  func(*DescribeFpgaImagesInput) DescribeFpgaImagesRequest
}

// Send marshals and sends the DescribeFpgaImages API request.
func (r DescribeFpgaImagesRequest) Send(ctx context.Context) (*DescribeFpgaImagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFpgaImagesResponse{
		DescribeFpgaImagesOutput: r.Request.Data.(*DescribeFpgaImagesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeFpgaImagesRequestPaginator returns a paginator for DescribeFpgaImages.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeFpgaImagesRequest(input)
//   p := ec2.NewDescribeFpgaImagesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeFpgaImagesPaginator(req DescribeFpgaImagesRequest) DescribeFpgaImagesPaginator {
	return DescribeFpgaImagesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeFpgaImagesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeFpgaImagesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeFpgaImagesPaginator struct {
	aws.Pager
}

func (p *DescribeFpgaImagesPaginator) CurrentPage() *DescribeFpgaImagesOutput {
	return p.Pager.CurrentPage().(*DescribeFpgaImagesOutput)
}

// DescribeFpgaImagesResponse is the response type for the
// DescribeFpgaImages API operation.
type DescribeFpgaImagesResponse struct {
	*DescribeFpgaImagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFpgaImages request.
func (r *DescribeFpgaImagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}