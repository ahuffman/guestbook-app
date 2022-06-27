/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apps

import (
	v1apps "guestbook-app/apis/apps/v1"
	v1guestbook "guestbook-app/apis/apps/v1/guestbook"
)

// Code generated by operator-builder. DO NOT EDIT.

// GuestbookLatestGroupVersion returns the latest group version object associated with this
// particular kind.
var GuestbookLatestGroupVersion = v1apps.GroupVersion

// GuestbookLatestSample returns the latest sample manifest associated with this
// particular kind.
var GuestbookLatestSample = v1guestbook.Sample(false)
