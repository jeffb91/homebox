package services

import (
	"context"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sysadminsmedia/homebox/backend/internal/data/repo"
)

func TestItemService_AddAttachment(t *testing.T) {
	temp := os.TempDir()

	svc := &ItemService{
		repo:     tRepos,
		filepath: temp,
	}

	attachmentSvc := &AttachmentService{
		repo: tRepos, // of wat de AttachmentService constructor vereist
	}

	loc, err := tRepos.Locations.Create(context.Background(), tGroup.ID, repo.LocationCreate{
		Description: "test",
		Name:        "test",
	})
	require.NoError(t, err)
	assert.NotNil(t, loc)

	itmC := repo.ItemCreate{
		Name:        fk.Str(10),
		Description: fk.Str(10),
		LocationID:  loc.ID,
	}

	itm, err := svc.repo.Items.Create(context.Background(), tGroup.ID, itmC)
	require.NoError(t, err)
	assert.NotNil(t, itm)
	t.Cleanup(func() {
		err := svc.repo.Items.Delete(context.Background(), itm.ID)
		require.NoError(t, err)
	})

	contents := fk.Str(1000)
	reader := strings.NewReader(contents)

	// Setup
	//afterAttachment, err := attachmentSvc.AttachmentAdd(tCtx, itm.ID, "testfile.txt", "attachment", false, reader)
	attachment, err := attachmentSvc.AttachmentAdd(tCtx, repo.AttachmentCreate{
		RelatedType: "item", // of wat je voor items gebruikt
		RelatedID:   itm.ID,
		Title:       "testfile.txt",
		Type:        "attachment",
		File:        reader,
		Primary:     false,
	})
	require.NoError(t, err)
	//assert.NotNil(t, afterAttachment)
	assert.NotNil(t, attachment)

	// Check that the file exists
	//storedPath := afterAttachment.Attachments[0].Path
	storedPath := attachment.Path

	// {root}/{group}/{item}/{attachment}
	//assert.Equal(t, path.Join(temp, "homebox", tGroup.ID.String(), "documents"), path.Dir(storedPath))
	expectedDir := path.Join(temp, "item", itm.ID.String(), "documents")
	assert.Equal(t, expectedDir, path.Dir(storedPath))

	// Check that the file contents are correct
	bts, err := os.ReadFile(storedPath)
	require.NoError(t, err)
	assert.Equal(t, contents, string(bts))

	// Cleanup - verwijder het attachment
	t.Cleanup(func() {
		err := tRepos.Attachments.Delete(context.Background(), attachment.ID)
		require.NoError(t, err)
	})
}
